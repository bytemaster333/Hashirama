# 📥 Installation Guide for Hashirama Operator

This guide provides step-by-step instructions on how to install, verify, and use the Hashirama Kubernetes Operator.

## ✅ Prerequisites

Before you begin, ensure you have the following installed:

*   **Kubernetes Cluster**: You can use a local cluster like [Kind](https://kind.sigs.k8s.io/) or Minikube, or a remote cloud cluster.
*   **kubectl**: The Kubernetes command-line tool. [Install instructions](https://kubernetes.io/docs/tasks/tools/).

---

## 🚀 Option 1: Quick Install (Recommended)

You can install the operator directly from the internet without downloading the repository.

1.  **Run the install command**:

    ```sh
    kubectl apply -f https://raw.githubusercontent.com/bytemaster333/Hashirama/main/Hashirama/dist/install.yaml
    ```

    *This command installs all necessary components: Namespace, Custom Resource Definitions (CRDs), RBAC Roles, and the Controller Manager.*

2.  **Verify Installation**:

    Check if the operator pod is running:

    ```sh
    kubectl get pods -n hashirama-system
    ```

    You should see a pod named `hashirama-controller-manager-xxx` with status `Running`.

---

## 📦 Option 2: Install from Source

If you prefer to download the code first or want to inspect the files:

1.  **Clone the Repository**:

    ```sh
    git clone https://github.com/bytemaster333/Hashirama.git
    cd Hashirama
    ```

2.  **Install the Operator**:

    ```sh
    kubectl apply -f Hashirama/dist/install.yaml
    ```

3.  **Verify Installation**:

    ```sh
    kubectl get pods -n hashirama-system
    ```

---

## 🛠 Usage: Creating a MadaraChain

Once installed, you can deploy your own `MadaraChain`.

1.  **Create a YAML file** (e.g., `my-chain.yaml`):

    ```yaml
    apiVersion: batch.starknet.l3/v1alpha1
    kind: MadaraChain
    metadata:
      name: my-first-chain
    spec:
      chainID: "SN_L3_DEMO"
      replicas: 1
      port: 9944
      image: "ghcr.io/madara-alliance/madara:latest" 
    ```

2.  **Apply the configuration**:

    ```sh
    kubectl apply -f my-chain.yaml
    ```

3.  **Check the Chain status**:

    ```sh
    kubectl get madarachains
    ```

---

## 🗑️ Uninstall

To remove the operator and all related resources from your cluster:

```sh
kubectl delete -f https://raw.githubusercontent.com/bytemaster333/Hashirama/main/Hashirama/dist/install.yaml
```

*(Or if you cloned the repo: `kubectl delete -f Hashirama/dist/install.yaml`)*
