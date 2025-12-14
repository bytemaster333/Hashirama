# 📥 Installation Guide for Hashirama Operator

This guide provides step-by-step instructions on how to install, verify, and use the **Hashirama Kubernetes Operator** and its **Dashboard**.

## ✅ Prerequisites

Before you begin, ensure you have the following installed:

*   [**Git**](https://git-scm.com/)
*   [**Docker**](https://www.docker.com/) (Required for Kind)
*   [**Kind**](https://kind.sigs.k8s.io/) (To run the local cluster)
*   [**kubectl**](https://kubernetes.io/docs/tasks/tools/) (To interact with the cluster)
*   [**Node.js & npm**](https://nodejs.org/) (Required for the Dashboard)

---

## �️ Step-by-Step Installation

### 1. 📥 Clone the Repository

First, download the project code to your local machine.

```sh
git clone https://github.com/bytemaster333/Hashirama.git
cd Hashirama
```

### 2. 🚀 Create the Cluster

Start your local Kubernetes cluster using Kind. **This must be done first.**

```sh
kind create cluster --name hashirama
```

### 3. 📦 Install the Operator

Apply the installation manifest to your cluster.

```sh
kubectl apply -f Hashirama/dist/install.yaml
```

*(This installs the Namespace, CRDs, Roles, and the Controller Manager)*

### 4. 🖥️ Activate the Dashboard

The dashboard provides a web interface to manage your chains.

1.  Navigate to the dashboard directory:
    ```sh
    cd dashboard
    ```

2.  Install dependencies:
    ```sh
    npm install
    ```

3.  Start the development server:
    ```sh
    npm run dev
    ```

4.  **Open your browser** to: [http://localhost:3000](http://localhost:3000)

---

## 🎮 Usage

You can now use the Dashboard to create chains visually, or use `kubectl` manually.

### Creating a Chain via CLI

1.  **Create a file** `my-chain.yaml`:

    ```yaml
    apiVersion: batch.starknet.l3/v1alpha1
    kind: MadaraChain
    metadata:
      name: demo-chain
    spec:
      chainID: "SN_L3_001"
      replicas: 1
      port: 9944
    ```

2.  **Apply it**:

    ```sh
    kubectl apply -f my-chain.yaml
    ```

3.  **Check Status**:

    ```sh
    kubectl get madarachains
    ```

---

## 🗑️ Cleanup

To stop everything:

1.  **Delete the cluster**:
    ```sh
    kind delete cluster --name hashirama
    ```

2.  **Stop the dashboard**: Press `Ctrl + C` in your terminal.
