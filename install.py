
"""
Build and install KeyNuker to OpenWhisk
"""

import subprocess
import os
import collections
import shutil

def main():

    # Builds go binaries and packages into zip file
    build_binaries()

    # Installs openwhisk actions via wsk utility to your configured OpenWhisk system
    install_openwhisk_actions(get_default_packaging_params())

    print("Success!")

def build_binaries():
    """
    Recursively process all directories in project, and every folder that has a main.go:

    - Build go binaries
    - Package binaries into zip file
    """
    for path in dirs_with_main():
        print("Building action binary in path: {}".format(path))
        build_binary_in_path(path)

def build_binary_in_path(path):
    
    # Save the current working directory
    cwd = os.getcwd()
    
    os.chdir(path)

    go_build_main()
    zip_binary_main()

    # Restore the original current working directory
    os.chdir(cwd) 
    
def go_build_main():
    """
    Build the main.go file into an "exec" binary 
    """
    assert_go_version()

    # Build the action binary 
    subprocess.check_call("env GOOS=linux GOARCH=amd64 go build -o exec main.go", shell=True)
 
def zip_binary_main():
    """
    Bundle the binary into a zip file
    """

    # Create the zip file
    subprocess.check_call("zip action.zip exec", shell=True)

def assert_go_version():
    """
    Make sure the go version is new enough.
    """
    goversion = subprocess.check_output(["go", "version"])
    if "go1.8" not in goversion and "go1.9" not in goversion:
        raise Exception("Your go version is too old.  Must use go 1.8 or later" + 
            " due to https://groups.google.com/d/msg/golang-nuts/9SaVxumSc-Y/rNAI8R7_BAAJ") 
    
def dirs_with_main():

    result = []

    for root, dirs, files in os.walk("cmd"):
        if "main.go" in files:
            result.append(root)

    return result

def install_openwhisk_actions(packaging_params):

    # Each action definition needs to be declared with certain default parameters (like a config binding)
    # And those parameters are set based on the environment variables declared in docs/install.adoc.  
    # This map defines all of the required params for each action, and which env variable it should get
    # the value from 
    action_params_to_env = {
        "fetch-aws-keys":{
            "TargetAwsAccounts": "TARGET_AWS_ACCOUNTS",
            "KeyNukerOrg": "KEYNUKER_ORG",
        },
        "github-user-aggregator": {
            "GithubAccessToken": "GITHUB_ACCESS_TOKEN",
            "GithubOrgs": "GITHUB_ORGS",
            "KeyNukerOrg": "KEYNUKER_ORG",            
        },
        "github-user-events-scanner": {
            "GithubAccessToken": "GITHUB_ACCESS_TOKEN",
        },
        "lookup-github-users-aws-keys": {
            "username": "KEYNUKER_DB_KEY",
            "host": "KEYNUKER_DB_HOST",
            "password": "KEYNUKER_DB_SECRET_KEY",
            "dbname": "KEYNUKER_DB_NAME",
        },
        "nuke-leaked-aws-keys": {
            "TargetAwsAccounts": "TARGET_AWS_ACCOUNTS",
        },
        "write-doc": {
            "username": "KEYNUKER_DB_KEY",
            "host": "KEYNUKER_DB_HOST",
            "password": "KEYNUKER_DB_SECRET_KEY",
            "dbname": "KEYNUKER_DB_NAME",
        },
        "monitor-activations": {},
    }

    actions = []
    for path in dirs_with_main():
        print("Installing OpenWhisk action for path: {}".format(path))
        packaging_params.path = path
        action = install_openwhisk_action_in_path(packaging_params, action_params_to_env, path)
        actions.append(action)

    sequences = install_openwhisk_action_sequences(actions)

    install_openwhisk_alarm_triggers()
    
    install_openwhisk_rules(sequences, actions)


def install_openwhisk_action_sequences(available_actions):
    """
    Create a list of action sequences.  The command line equivalent of:

    $ wsk action create fetch-aws-keys-write-doc --sequence fetch-aws-keys,write-doc
    """

    # Dictionary of action sequences
    # Key: name of action sequence
    # Value: ordered list of actions that comprise the sequence
    action_sequences = {
        "fetch-aws-keys-write-doc": [
            "fetch-aws-keys",
            "write-doc",
        ],
        "github-user-aggregator-write-doc": [
            "github-user-aggregator",
            "write-doc",
        ],
        "github-user-events-scanner-nuker": [
            "lookup-github-users-aws-keys",
            "github-user-events-scanner",
            "nuke-leaked-aws-keys",
        ]
    }

    for action_sequence, actions in action_sequences.iteritems():

        # Make sure all the actions in this sequence are valid.  This protects
        # against bugs due to renaming actions, and forgetting to update the action_sequences dictionary
        for action in actions:
            if action not in available_actions:
                raise Exception("Cannot create action sequence that contains invalid action: {}".format(action))

        # If the action sequence already exists, delete it
        if openwhisk_action_exists(action_sequence):
            delete_openwhisk_action(action_sequence)

        # Get the actions list as a comma delimed string, eg: fetch-aws-keys,write-doc
        comma_delimited_actions = ",".join(actions)

        # Default the action timeout to 5 minutes, which is the max value on the hosted IBM bluemix platform
        command = "wsk action create {} --timeout 300000 --sequence {}".format(
            action_sequence,
            comma_delimited_actions,
        )
        
        subprocess.check_call(command, shell=True)


    return action_sequences.keys()


def install_openwhisk_action_in_path(packaging_params, action_params_to_env, path):

    """
    This performs the equivalent of the command line:

    wsk action create fetch_aws_keys --docker tleyden5iwx/openwhisk-dockerskeleton --param AwsAccessKeyId "$AWS_ACCESS_KEY_ID" --param AwsSecretAccessKey "$AWS_SECRET_ACCESS_KEY" --param KeyNukerOrg "default"
    
    Where the param values are pulled out of environment variables.  The param vals and their
    corresponding environment variable names are specified in the action_params_to_env dictionary

    The name of the action is discovered from the path basename.
    """

    # Save the current working directory
    cwd = os.getcwd()
    
    os.chdir(path)

    openwhisk_action = os.path.basename(os.getcwd())
    params_to_env = action_params_to_env[openwhisk_action]

    if not openwhisk_action_exists(openwhisk_action):
        install_openwhisk_action(
            packaging_params,
            openwhisk_action,
            params_to_env,
            )
    else:
        delete_openwhisk_action(openwhisk_action)
        install_openwhisk_action(
            packaging_params,
            openwhisk_action,
            params_to_env,
            )


    # Restore the original current working directory
    os.chdir(cwd) 

    return openwhisk_action

def install_openwhisk_alarm_triggers():
    """
    This installs the following triggers:

    $ wsk trigger create every4Hours --feed /whisk.system/alarms/alarm -p cron '0 */4 * * *'
    $ wsk trigger create every15Minutes --feed /whisk.system/alarms/alarm -p cron '*/15 * * * *'
    """
    alarm_triggers = {
        "every4Hours": "0 */4 * * *",
        "every15Minutes": "*/15 * * * *",
    }
    for alarm_trigger, schedule in alarm_triggers.iteritems():

        if openwhisk_trigger_exists(alarm_trigger):
            delete_openwhisk_trigger(alarm_trigger)
        
        command = "wsk trigger create {} --feed /whisk.system/alarms/alarm --param cron '{}'".format(
            alarm_trigger,
            schedule,
        )
        
        subprocess.check_call(command, shell=True)

def install_openwhisk_rules(available_sequences, available_actions):
    """
    This installs a rule that connects a trigger (an alarm in our case) to an action.
    For examle, this will cause the fetch-aws-keys-write-doc action to run every four hours:

    $ wsk rule create scheduled-fetch-aws-keys-write-doc every4Hours fetch-aws-keys-write-doc
    """

    rules = {
        "scheduled-fetch-aws-keys-write-doc": {
            "trigger": "every4Hours", 
            "action": "fetch-aws-keys-write-doc",
        },
        "scheduled-github-user-aggregator-write-doc": {
            "trigger": "every4Hours", 
            "action": "github-user-aggregator-write-doc",
        },
        "scheduled-github-user-events-scanner-nuker": {
            "trigger": "every15Minutes", 
            "action": "github-user-events-scanner-nuker",
        },
    }

    for rule, rule_target in rules.iteritems():
        trigger = rule_target["trigger"]
        action = rule_target["action"]

        # Make sure all the actions in this sequence are valid.  This protects
        # against bugs due to renaming actions, and forgetting to update the action_sequences dictionary
        if action not in available_actions and action not in available_sequences:
            raise Exception("Invalid action: {}.  Not in available_actions: {} or available_sequences: {}".format(action, available_actions, available_sequences))

        if openwhisk_rule_exists(rule):
            delete_openwhisk_rule(rule)
        
        command = "wsk rule create {} {} {}".format(
            rule,
            trigger,
            action,
        )
        
        subprocess.check_call(command, shell=True)


def delete_openwhisk_action(openwhisk_action):
    command = "wsk action delete {}".format(
        openwhisk_action,
    )
    subprocess.check_call(command, shell=True)

def delete_openwhisk_trigger(openwhisk_trigger):
    command = "wsk trigger delete {}".format(
        openwhisk_trigger,
    )
    subprocess.check_call(command, shell=True)

def delete_openwhisk_rule(openwhisk_rule):
    command = "wsk rule delete {}".format(
        openwhisk_rule,
    )
    subprocess.check_call(command, shell=True)


def install_openwhisk_action(packaging_params, openwhisk_action, params_to_env):

    expanded_params = expand_params(params_to_env)

    if packaging_params.useDockerSkeleton == True:
        # Default the action timeout to 5 minutes, which is the max value on the hosted IBM bluemix platform
        command = "wsk action create {} --timeout 300000 --docker tleyden5iwx/openwhisk-dockerskeleton action.zip {}".format(
            openwhisk_action,
            expanded_params,
        )
    else:

        build_docker_in_path(packaging_params.path)

        # Default the action timeout to 5 minutes, which is the max value on the hosted IBM bluemix platform
        command = "wsk action create {} --timeout 300000 --docker {}/{} {}".format(
            openwhisk_action,
            discover_dockerhub_user(),
            openwhisk_action,
            expanded_params,
        )

    print("Installing OpenWhisk action via {}".format(command))
    subprocess.check_call(command, shell=True)

def expand_params(params_to_env):
    """
    Given a dictionary like:

    {
        "AwsAccessKeyId": "AWS_ACCESS_KEY_ID",
        "AwsSecretAccessKey": "AWS_SECRET_ACCESS_KEY",
        "KeyNukerOrg": "KEYNUKER_ORG",
    }

    Convert to a string like:

    --param KeyNukerOrg default --param AwsAccessKeyId ***** --param AwsSecretAccessKey ********

    Where the param values are created based on the contents of the corresponding environment
    variable (eg, AwsAccessKeyId)

    """

    result_list = []

    for paramName, envVarName in params_to_env.iteritems():
        if paramName == "GithubOrgs":
            # This needs special handling since it's an array
            continue
        if paramName == "TargetAwsAccounts":
            # This needs special handling since it's an array
            continue

        envVarVal = os.environ.get(envVarName)
        if envVarVal is None:
            raise Exception("You must set the {} environment variable".format(envVarName))
        result_list.append("--param")
        result_list.append(paramName)
        result_list.append('{}'.format(envVarVal))

    result = " ".join(result_list)

    if "GithubOrgs" in params_to_env:
        envVarName = params_to_env["GithubOrgs"]
        envVarVal = os.environ.get(envVarName)
        result += " --param GithubOrgs "
        result += "\'{}\'".format(envVarVal)

    if "TargetAwsAccounts" in params_to_env:
        envVarName = params_to_env["TargetAwsAccounts"]
        envVarVal = os.environ.get(envVarName)
        result += " --param TargetAwsAccounts "
        result += "\'{}\'".format(envVarVal)

    return result 


def update_openwhisk_action(openwhisk_action, params_to_env):
    raise Exception("Not implemented")

def openwhisk_action_exists(openwhisk_action):
    command = "wsk action get {} parameters".format(
        openwhisk_action,
    )
    return subprocess.call(command, shell=True) == 0

def openwhisk_trigger_exists(openwhisk_trigger):
    # TODO: detect if it exists.  
    return True

def openwhisk_rule_exists(openwhisk_rule):
    # TODO: detect if it exists.  
    return True

def build_docker_in_path(path):

    docker_build()
    docker_push()


def docker_build():
    """
    Generate and run a command like:
    docker build -t youruser/fetch-aws-keys .

    - The dockerhub user will be discovered from an environment variable
    - The dockerhub repo name will be disovered from the last path component of the current directory
    """

    dockerhub_user = discover_dockerhub_user()
    dockerhub_repo = discover_dockerhub_repo()

    # the Dockerfile lives in the parent directory of the current directory, so copy it into the
    # current directory
    cwd = os.getcwd()
    parent = os.path.dirname(cwd)
    src_dockerfile = os.path.join(parent, "Dockerfile")
    dest_dockerfile = "Dockerfile"
    shutil.copyfile(src_dockerfile, dest_dockerfile)

    # Build docker image
    subprocess.check_call("docker build -t {}/{} .".format(dockerhub_user, dockerhub_repo), shell=True)

    # Delete the Dockerfile copy that is no longer needed
    os.remove(dest_dockerfile)

def docker_push():
    """
    Generate and run a command like:
    docker push youruser/fetch-aws-keys

    - The dockerhub user will be discovered from an environment variable
    - The dockerhub repo name will be disovered from the last path component of the current directory
    """

    dockerhub_user = discover_dockerhub_user()
    dockerhub_repo = discover_dockerhub_repo()

    subprocess.check_call("docker push {}/{}".format(dockerhub_user, dockerhub_repo), shell=True)


def discover_dockerhub_user():
    dockerhub_user = os.environ.get("DOCKERHUB_USERNAME")
    if dockerhub_user is None:
        raise "You must set the DOCKERHUB_USERNAME environment variable"
    return dockerhub_user

def discover_dockerhub_repo():
    # If we are in the /home/go/src/github.com/tleyden/keynuker/cmd/fetch-aws-keys directory,
    # return the basename (fetch-aws-keys) which will be used to derive the name for the
    # dockerhub repo to push to
    return os.path.basename(os.getcwd())

def get_default_packaging_params():

    # Parameters to specify how the openwhisk actions are packaged
    packaging_params = collections.namedtuple('PackagingParams', 'useDockerSkeleton', 'path')

    # useDockerSkeleton: true or false.  True to use https://hub.docker.com/r/tleyden5iwx/openwhisk-dockerskeleton/
    #                                    False to directly build an image and push to dockerhub
    # There are two reasons you might want to set this to False:
    #   1. Want full control of all the code, as opposed to trusting the code in https://hub.docker.com/r/tleyden5iwx/openwhisk-dockerskeleton/
    #   2. Suspect there is an issue with the actionproxy.py wrapper code in openwhisk-dockerskeleton, and want to compare behavior.
    # If you set to False, you will need to have docker locally installed and a few extra environment
    # variables set.  You will also need to go into the cmd entrypoints and call "ow.RegisterAction(OpenWhiskCallback)"
    # rather than "keynuker_go_common.InvokeActionStdIo(OpenWhiskCallback)".
    packaging_params.useDockerSkeleton = True

    return packaging_params


if __name__ == "__main__":
    main() 
