# Multi-Service Cloud Architecture with Docker Swarm

A production-ready microservices architecture featuring 4 polyglot services deployed across a multi-node Azure VM Docker Swarm cluster (1 Manager, 2 Workers) with cross-region private VNet peering.

---

## Architecture Overview

* **Laravel Service** (`:8000`): PHP 8.4 backend utilizing SQLite with automated startup migrations, permission enforcement, and key generation.
* **Next.js Service** (`:3000`): Frontend client interface.
* **Node.js Service** (`:3001`): Microservice handling real-time APIs.
* **Go Service** (`:8080`): High-performance compiled backend worker.
* **Infrastructure**: Microsoft Azure Virtual Networks (VNets), bidirectional Global VNet Peering, and Docker Swarm overlay routing.

---

## Quick Start (Local Docker Compose)

To run the complete stack locally with zero configuration:

```bash
# Clone the repository
git clone [https://github.com/](https://github.com/)<YOUR-GITHUB-USERNAME>/<REPO-NAME>.git
cd <REPO-NAME>

# Start all containers
docker compose up -d --build