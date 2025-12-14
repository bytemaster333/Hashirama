# 📥 Installation Guide

Follow these steps to set up the Hashirama Operator and Dashboard for local development.

## 1. Create Cluster

Create a local Kubernetes cluster using Kind (if it doesn't exist).

```sh
kind create cluster --name hashirama
```

## 2. Install CRDs

Navigate to the operator directory and install the Custom Resource Definitions into your cluster.

```sh
cd Hashirama
make install
```

## 3. Run Controller Locally

Run the controller on your host machine. This allows for faster development cycles.

```sh
make run
```

*Keep this terminal window open to keep the controller running.*

## 4. Deploy Sample

Open a new terminal. You can edit the sample file `config/samples/batch_v1alpha1_madarachain.yaml` or use the following configuration:

```yaml
apiVersion: batch.starknet.l3/v1alpha1
kind: MadaraChain
metadata:
  name: madarachain-sample
spec:
  chainID: "my-l3-chain"
  replicas: 1
  port: 9944
```

Apply it to the cluster:

```sh
kubectl apply -f config/samples/batch_v1alpha1_madarachain.yaml
```

> [!NOTE]
> **ARM64 Workaround**: The controller is currently configured to use `nginx:latest` as a placeholder because the official Madara image does not support ARM64. To revert to the official image, modify `internal/controller/madarachain_controller.go` to remove the hardcoded nginx image and restore the original arguments.

## 5. Dashboard

I have created a Next.js dashboard to manage your chains visually.

### Run Dashboard:

1.  Navigate to the dashboard directory:
    ```sh
    cd ../dashboard
    ```
    *(Assuming you are still in `Hashirama` directory)*

2.  Install and Run:
    ```sh
    npm install
    npm run dev
    ```

### Access:

Open [http://localhost:3000](http://localhost:3000) in your browser.

- You can view existing chains.
- Click **"Create Chain"** to deploy a new one.
