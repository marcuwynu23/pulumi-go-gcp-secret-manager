package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/secretmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		projectID := config.Require(ctx, "gcp:project")
		region := config.Get(ctx, "gcp:region")
		if region == "" {
			region = "us-central1"
		}

		secretID := config.Get(ctx, "secretmanager:secretId")
		if secretID == "" {
			secretID = "my-api-key"
		}

		secretData := config.RequireSecret(ctx, "secretmanager:secretData")

		// Create a secret
		secret, err := secretmanager.NewSecret(ctx, "my-secret", &secretmanager.SecretArgs{
			SecretId: pulumi.String(secretID),
			Project:  pulumi.String(projectID),
			Replication: &secretmanager.SecretReplicationArgs{
				Auto: &secretmanager.SecretReplicationAutoArgs{},
			},
		})
		if err != nil {
			return err
		}

		// Add a secret version with the actual secret data
		secretVersion, err := secretmanager.NewSecretVersion(ctx, "my-secret-version", &secretmanager.SecretVersionArgs{
			Secret:     secret.ID(),
			SecretData: secretData,
		})
		if err != nil {
			return err
		}

		// Export outputs
		ctx.Export("secretId", secret.SecretId)
		ctx.Export("secretName", secret.Name)
		ctx.Export("secretVersion", secretVersion.Version)
		ctx.Export("secretData", pulumi.ToSecret(secretVersion.SecretData).(pulumi.StringOutput))

		return nil
	})
}
