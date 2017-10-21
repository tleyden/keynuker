#!/usr/bin/env bash

# make sure AWS_PROFILE is set
if [ -z "$AWS_PROFILE" ]; then
    echo "You must set the AWS_PROFILE environment variable to the aws config profile you want to use.  Eg, export AWS_PROFILE=DefaultAwsAccount"
    exit 1
fi

# make sure InitiatingAcccountAWSID is set
if [ -z "$InitiatingAcccountAWSID" ]; then
    echo "You must set the InitiatingAcccountAWSID environment variable"
    exit 1
fi

# make sure IAMRoleExternalID is set
if [ -z "$IAMRoleExternalID" ]; then
    echo "You must set the IAMRoleExternalID environment variable"
    exit 1
fi

wget https://raw.githubusercontent.com/tleyden/keynuker/master/docs/cloudformation/keynuker-allow-sts-assume-role.yml
echo "Installing Cloudformation Template with IAM Role with policy:"
cat keynuker-allow-sts-assume-role.yml

aws cloudformation create-stack --stack-name KeyNukerIAMRoleAccess --parameters ParameterKey=InitiatingAcccountAWSID,ParameterValue=$InitiatingAcccountAWSID ParameterKey=IAMRoleExternalID,ParameterValue=$IAMRoleExternalID  --template-body file://keynuker-allow-sts-assume-role.yml --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM --profile $AWS_PROFILE

echo "Waiting for cloudformation (I'll sleep for 60 seconds -- you go check your email and we'll see you back here in two hours)..."

sleep 60

echo "Success!  You now have a Cloudformation named KeyNukerIAMRoleAccess and a CrossAccount role with an associated policy that allows ListUsers, ListKeys, DeleteKeys.  To uninstall, delete the KeyNukerIAMAccess Cloudformation from the Web UI or via the CLI"