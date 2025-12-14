#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Application: Hashirama (Starknet Operator)${NC}"
echo -e "${BLUE}🔧 Starting One-Click Local Development Setup...${NC}\n"

# 1. Prerequisites Check
echo -e "${YELLOW}🔍 Checking Prerequisites...${NC}"

command_exists() {
    command -v "$1" >/dev/null 2>&1
}

if ! command_exists docker; then
    echo -e "${RED}❌ Error: Docker is not installed.${NC}"
    exit 1
fi

if ! command_exists kind; then
    echo -e "${RED}❌ Error: Kind is not installed.${NC}"
    exit 1
fi

if ! command_exists kubectl; then
    echo -e "${RED}❌ Error: Kubectl is not installed.${NC}"
    exit 1
fi

if ! command_exists node; then
    echo -e "${RED}❌ Error: Node.js is not installed.${NC}"
    exit 1
fi

if ! command_exists npm; then
    echo -e "${RED}❌ Error: npm is not installed.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ All prerequisites found!${NC}\n"

# 2. Cluster Management
CLUSTER_NAME="hashirama"
echo -e "${YELLOW}☸️  Checking Kind Cluster '${CLUSTER_NAME}'...${NC}"

if kind get clusters | grep -q "^${CLUSTER_NAME}$"; then
    echo -e "${GREEN}✅ Cluster '${CLUSTER_NAME}' already exists.${NC}"
    echo -e "🔄 Switching context to 'kind-${CLUSTER_NAME}'..."
    kubectl config use-context "kind-${CLUSTER_NAME}"
else
    echo -e "${BLUE}🆕 Creating Kind cluster '${CLUSTER_NAME}'...${NC}"
    kind create cluster --name "${CLUSTER_NAME}"
    echo -e "${GREEN}✅ Cluster created successfully!${NC}"
fi
echo ""

# 3. Installation
echo -e "${YELLOW}📦 Installing CRDs (make install)...${NC}"
make install
echo -e "${GREEN}✅ CRDs installed!${NC}\n"

# 4. Deployment (The Trick)
SAMPLE_FILE="config/samples/batch_v1alpha1_madarachain.yaml"
echo -e "${YELLOW}⚡ Deploying Sample MadaraChain (declarative pre-creation)...${NC}"
if [ -f "$SAMPLE_FILE" ]; then
    kubectl apply -f "$SAMPLE_FILE"
    echo -e "${GREEN}✅ Sample Request Applied!${NC} (The controller will pick this up immediately upon start)"
else
    echo -e "${RED}⚠️  Warning: Sample file '$SAMPLE_FILE' not found. Skipping sample deployment.${NC}"
fi
echo ""

# 5. Dashboard Setup (Automatic)
echo -e "${YELLOW}📊 Setting up Dashboard...${NC}"
# Save current directory (root of Hashirama)
REPO_ROOT=$(pwd)

# Dashboard is expected to be a sibling folder
if [ -d "../dashboard" ]; then
    cd ../dashboard
else
    echo -e "${RED}❌ Error: '../dashboard' directory not found!${NC}"
    echo -e "Expected structure: \n  - Hashirama/ (contains setup.sh)\n  - dashboard/"
    exit 1
fi

# Check if node_modules exists to skip install if possible (or just run npm install which is smart enough)
if [ ! -d "node_modules" ]; then
    echo -e "${BLUE}📦 Installing dashboard dependencies (npm install)...${NC}"
    npm install
else
    echo -e "${GREEN}✅ Dependencies already installed.${NC}"
fi

echo -e "${BLUE}🚀 Starting Dashboard in background...${NC}"
# Run in background, redirect logs to prevent clutter
# Log file placed back in Hashirama folder for visibility
npm run dev > "$REPO_ROOT/dashboard.log" 2>&1 &
DASHBOARD_PID=$!

# Define cleanup function
cleanup() {
    echo -e "\n${YELLOW}🧹 Shutting down...${NC}"
    echo -e "${BLUE}Stopping Dashboard (PID: $DASHBOARD_PID)...${NC}"
    # Check if process is still running before killing
    if ps -p $DASHBOARD_PID > /dev/null; then
        kill $DASHBOARD_PID
    fi
    echo -e "${GREEN}✅ Done.${NC}"
}

# Trap EXIT to ensure dashboard is killed when script exits
trap cleanup EXIT
# Go back to Hashirama folder for make run
cd "$REPO_ROOT"

echo -e "${GREEN}✅ Dashboard running at: http://localhost:3000${NC}"
echo -e "(Dashboard logs are being written to dashboard.log)\n"

# 6. Start Controller
echo -e "${YELLOW}🎮 Starting Controller (make run)...${NC}"
echo -e "${BLUE}Logs will appear below. Press Ctrl+C to stop everything.${NC}"
echo "---------------------------------------------------"
# Use direct call (not exec) so trap works
make run
