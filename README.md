# VCDeploy

**Keep in mind, this project is currently in the pre-release phase. This tool only works with Vapor Cloud 2**

VCDeploy provides a simple binary, making it possible to trigger a deploy from for example a CI system. The binary works on all Linux distributions.

When you trigger a deploy, it will give all the output in real-time, so the status is posted through the CI system.

## Usage

Start by downloading the binary to the system.

```
curl -O https://github.com/qutheory/vc-deploy/releases/download/v0.1.0/vc-deploy
chmod +x vc-deploy
```

Now you can run it:

```
./vc-deploy --app <app-slug> --env <env> --token <developer-token> 
```

**Important:** The developer token, should be kept private, as it gives full access to the Vapor Cloud 2 API.

If you want to set a custom branch to deploy, you can specify it like this. If you don't specify a branch, the environments default branch will be used.

```
./vc-deploy --app my-app --env my-env --token my-developer-token --branch my-branch
```
