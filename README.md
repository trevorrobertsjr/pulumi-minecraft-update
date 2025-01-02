# pulumi-minecraft-update

## How to use this repo:
Check out the [blog post](https://www.trevorrobertsjr.com/blog/minecraft-update-automation/) for more details.

Deploy this code with Pulumi

```bash
pulumi up
```

The Pulumi execution will output the value you should specify for the `document-name` parameter in the command below.

Get your Minecraft server's instance ID from the EC2 service to use with the InstanceId below.

Get the latest JAR download URL from the Minecraft web site to use with the MinecraftJarUrl parameter in the command below.

Update the parameters in this command and execute it. You can monitor the execution in the AWS Systems Manager console.

```bash
aws ssm start-automation-execution \
    --document-name "Your-Document-Name" \
    --parameters "InstanceId=i-yourinstanceid,MinecraftJarUrl=https://example.com/path/to/server.jar"

```