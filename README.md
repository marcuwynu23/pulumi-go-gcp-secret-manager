# Pulumi GCP Go: Secret Manager

This project provisions GCP Secret Manager resources using Pulumi and Go. It demonstrates how to:
- Use the Pulumi GCP provider in a Go program
- Create secrets and secret versions
- Handle sensitive data with Pulumi secrets

## Providers

- Google Cloud Platform via the Pulumi GCP SDK for Go (`github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp`)

## Resources

- **Secret** (`gcp.secretmanager.Secret`)
  - Secret container with automatic replication

- **Secret Version** (`gcp.secretmanager.SecretVersion`)
  - Version with the actual secret data (stored as Pulumi secret)

## Outputs

- **secretId**: The ID of the created secret
- **secretName**: The full resource name of the secret
- **secretVersion**: The version of the created secret
- **secretData**: The secret data (encrypted in state)

## When to Use This Project

- You need to store API keys, passwords, or other sensitive data
- You want to manage secrets programmatically
- You need versioned secrets with automatic replication

## Prerequisites

- Go 1.21+ installed
- A Google Cloud account with billing enabled
- GCP credentials configured for Pulumi (via `gcloud auth application-default login`)
- Secret Manager API enabled

## Usage

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Configure your stack:
   ```bash
   cp Pulumi.dev.yaml.example Pulumi.dev.yaml
   pulumi config set gcp:project YOUR_PROJECT_ID
   pulumi config set secretmanager:secretId my-api-key
   pulumi config set --secret secretmanager:secretData YOUR_SECRET_VALUE
   ```

3. Preview and deploy:
   ```bash
   pulumi preview
   pulumi up
   ```

4. Access the secret in other programs:
   ```bash
   pulumi stack output secretData --show-secrets
   ```

## Project Layout

```
├── Pulumi.yaml                  Pulumi project definition
├── Pulumi.dev.yaml.example      Template for local dev configuration
├── go.mod                       Go module declaration and dependencies
├── main.go                      Pulumi program defining secret resources
├── .gitignore                   Git ignore rules
└── LICENSE                      MIT License
```

## Configuration

| Name | Description | Default |
|------|-------------|---------|
| `gcp:project` | The Google Cloud project to deploy into | _required_ |
| `gcp:region` | The GCP region for resources | `us-central1` |
| `secretmanager:secretId` | The ID of the secret to create | `my-api-key` |
| `secretmanager:secretData` | The secret data value (secret) | _required_ |

## Next Steps

- Add IAM bindings for secret access control
- Configure secret rotation policies
- Set up secret replication across regions
- Integrate with Cloud Run, Cloud Functions, or GKE

## Getting Help

- Pulumi Documentation: https://www.pulumi.com/docs/
- GCP Provider Reference: https://www.pulumi.com/registry/packages/gcp/
- Community Slack: https://slack.pulumi.com/
- GitHub Issues: https://github.com/pulumi/pulumi/issues