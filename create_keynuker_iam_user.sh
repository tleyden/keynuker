

# make sure AWS_PROFILE is set
if [ -z "$AWS_PROFILE" ]; then
    echo "You must set the AWS_PROFILE environment variable to the aws config profile you want to use.  Eg, export AWS_PROFILE=DefaultAwsAccount"
    exit 1
fi

wget https://raw.githubusercontent.com/tleyden/keynuker/master/docs/keynuker-iam-user-policy.yml
echo "Installing Cloudformation Template with IAM User + Policy + Access key"
cat keynuker-iam-user-policy.yml
aws cloudformation create-stack --stack-name KeyNukerIAMAccess --template-body file://keynuker-iam-user-policy.yml --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM --profile $AWS_PROFILE && \
aws cloudformation describe-stacks --stack-name KeyNukerIAMAccess --query 'Stacks[0].Outputs[?OutputKey==`AwsAccessKey`].OutputValue' --output text --profile $AWS_PROFILE && \
aws cloudformation describe-stacks --stack-name KeyNukerIAMAccess --query 'Stacks[0].Outputs[?OutputKey==`AwsSecretAccessKey`].OutputValue' --output text --profile $AWS_PROFILE

echo "Success!  You now have a Cloudformation Template named KeyNukerIAMAccess and an IAM user named KeyNuker with an associated policy that allows ListUsers, ListKeys, DeleteKeys.  To uninstall, delete the Cloudformation from the Web UI or via the CLI"