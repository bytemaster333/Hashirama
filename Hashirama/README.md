# 🏯 Hashirama: Kubernetes Operator for Starknet L3s

**Hashirama** is a powerful Kubernetes Operator designed to automate the deployment and management of **Starknet Layer 3 (L3)** appchains using [Madara](https://github.com/keep-starknet-strange/madara).

It simplifies the complexity of running a zk-rollup by providing a cloud-native, declarative API.

![Starknet](https://img.shields.io/badge/Starknet-L3-blue)
![Kubernetes](https://img.shields.io/badge/Kubernetes-Operator-326ce5)
![Status](https://img.shields.io/badge/Status-Alpha-orange)

---

## ✨ Features

*   **Declarative API**: Define your L3 chain as a simple Kubernetes Custom Resource (`MadaraChain`).
*   **One-Click Deployment**: Automates the creation of StatefulSets, Services, and storage configurations.
*   **Visual Dashboard**: Includes a built-in Next.js dashboard to manage your chains visually.
*   **Local Development Friendly**: Comes with a "One-Click" setup script for automated local testing with Kind.

---

## 🚀 Quick Start (Local Development)

We provide a specialized script that automates the entire setup process (Cluster creation, CRD installation, Dashboard setup, and Controller execution).

**Prerequisites:**
*   [Docker](https://www.docker.com/)
*   [Kind](https://kind.sigs.k8s.io/)
*   [kubectl](https://kubernetes.io/docs/tasks/tools/)
*   [Node.js](https://nodejs.org/) (for Dashboard)

**Run the Setup Script:**

```bash
./setup.sh
```

This script will:
1.  Check/Install the Kind cluster.
2.  Install the Operator.
3.  **Start the Dashboard** automatically at `http://localhost:3000`.
4.  Start the Controller and stream logs to your terminal.

> **⚠️ Developer Note (Apple Silicon Users):**
> By default, this operator deploys a lightweight `nginx` placeholder image. This is a strategic workaround strictly for local testing on Apple Silicon (M1/M2/M3) chips due to current upstream compatibility issues.
>
> **Please Note:** This project is fully architected for and compatible with standard open-source ecosystems (Ubuntu, Debian, Linux) and production cloud environments (AWS, GCP). For standard deployments, the operator natively supports the official `madara:latest` sequencer image without modification—simply update the `image` field in your CRD.

---

## 🛠 Manual Installation

If you prefer to install manually or are deploying to a shared cluster, please see our detailed [**Installation Guide (INSTALL.md)**](INSTALL.md).

---

## 🎮 Usage

You can create a Starknet L3 chain using `kubectl` or the Dashboard.

### Creating a Chain (YAML)

Create a file named `chain.yaml`:

```yaml
apiVersion: batch.starknet.l3/v1alpha1
kind: MadaraChain
metadata:
  name: my-starknet-appchain
spec:
  chainID: "SN_APPCHAIN_001"
  replicas: 1
  port: 9944
```

Apply it:

```bash
kubectl apply -f chain.yaml
```

The operator will immediately spin up the necessary pods and services.

---

## 🏗 Architecture

The project consists of two main components:

1.  **Controller Manager**: Included in `./cmd` and `./internal`. It watches for `MadaraChain` resources and reconciles the cluster state.
2.  **Dashboard**: Located in the [`dashboard/`](./dashboard) directory. A Next.js application that interacts with the Kubernetes API to provide a UI.

---

## 🤝 Contributing

Contributions are welcome!

1.  Fork the repository.
2.  Create your feature branch (`git checkout -b feature/amazing-feature`).
3.  Commit your changes (`git commit -m 'Add some amazing feature'`).
4.  Push to the branch (`git push origin feature/amazing-feature`).
5.  Open a Pull Request.

---

**Built with ❤️ for the Starknet Community.**
