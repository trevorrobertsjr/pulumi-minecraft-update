package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an SSM Document for automation
		documentContent := `{
		  "schemaVersion": "0.3",
		  "description": "Automate AMI creation and Minecraft server update",
		  "parameters": {
		    "InstanceId": {
		      "type": "String",
		      "description": "The ID of the instance to create an AMI for and update Minecraft."
		    },
		    "MinecraftJarUrl": {
		      "type": "String",
		      "description": "The URL of the Minecraft server.jar file."
		    }
		  },
		  "mainSteps": [
		    {
		      "name": "CreateAMI",
		      "action": "aws:createImage",
		      "inputs": {
		        "InstanceId": "{{ InstanceId }}",
		        "ImageName": "minecraft-{{ InstanceId }}-backup-{{ global:DATE_TIME }}"
		      },
		      "isCritical": true
		    },
		    {
		      "name": "UpdateMinecraftServer",
		      "action": "aws:runCommand",
		      "inputs": {
		        "DocumentName": "AWS-RunShellScript",
		        "InstanceIds": ["{{ InstanceId }}"],
		        "Parameters": {
		          "commands": [
		            "sudo su - minecraft -c 'sudo systemctl stop minecraft'",
		            "sudo su - minecraft -c 'curl -o /opt/minecraft/server/server.jar {{ MinecraftJarUrl }}'",
		            "sudo su - minecraft -c 'sudo systemctl start minecraft'"
		          ]
		        }
		      }
		    }
		  ]
		}`

		document, err := ssm.NewDocument(ctx, "UpdateMinecraftServer", &ssm.DocumentArgs{
			Content:      pulumi.String(documentContent),
			DocumentType: pulumi.String("Automation"),
		})
		if err != nil {
			return err
		}

		ctx.Export("DocumentName", document.Name)

		return nil
	})
}
