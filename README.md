# My Pipeline 163 — CI/CD для Go-приложения

[![GitLab CI](https://img.shields.io/badge/GitLab_CI-passing-brightgreen)](https://gitlab.com/vladimir_vgb/my-pipeline-163/-/pipelines)
[![Jenkins](https://img.shields.io/badge/Jenkins-passing-brightgreen)](http://localhost:8080/job/pipeline-from-gitlab-ssh/)

---

## 🇷🇺 Русская версия

### 📖 О проекте

Демонстрационный проект **полного цикла CI/CD** для Go-приложения:
от коммита до продакшена в Kubernetes. Проект показывает интеграцию
**GitLab CI** и **Jenkins**, работу с **Docker Hub**, **HashiCorp Vault**
и **Kubernetes** (Minikube).

### 🛠 Технологии

| Компонент | Технология |
|---|---|
| **Язык** | Go 1.21 |
| **CI/CD** | GitLab CI, Jenkins (Pipeline as Code) |
| **Контейнеризация** | Docker, Docker Hub |
| **Оркестрация** | Kubernetes (Minikube) |
| **Управление секретами** | HashiCorp Vault (JWT auth) |
| **Мониторинг** | Prometheus, Grafana |
| **GitOps** | Flux |

### 🚀 Как работает пайплайн

1. **`git push` в `main`** — GitLab получает коммит.
2. **GitLab CI** запускает пайплайн: `build → test → deploy`.
3. **GitLab webhook** отправляет событие в Jenkins через **localhost.run tunnel**.
4. **Jenkins**:
   - Checkout кода
   - Сборка Docker-образа в **Docker-агенте**
   - Push образа в **Docker Hub**
   - Деплой в **Kubernetes**
   - Проверка подов и сервиса
5. **Vault** выдаёт секреты для GitLab CI через **JWT-аутентификацию**.

### 📁 Структура проекта

```
my-pipeline-163/
├── main.go                 # Go-приложение
├── go.mod                  # Go-модули
├── Dockerfile              # Многостадийная сборка
├── Jenkinsfile             # Jenkins Pipeline
├── .gitlab-ci.yml          # GitLab CI конфигурация
├── k8s/                    # Kubernetes манифесты
│   ├── deployment.yaml
│   └── service.yaml
└── README.md
```

### 🎯 Быстрый старт

```bash
# 1. Клонировать репозиторий
git clone git@gitlab.com:vladimir_vgb/my-pipeline-163.git
cd my-pipeline-163

# 2. Собрать Docker-образ
docker build -t zakon13/my-go-app:latest .

# 3. Запустить локально
docker run -p 8080:8080 zakon13/my-go-app:latest

# 4. Запустить пайплайн
git push origin main
```

### 🔗 Ссылки

- **GitLab репозиторий**: https://gitlab.com/vladimir_vgb/my-pipeline-163
- **Docker Hub образ**: https://hub.docker.com/r/zakon13/my-go-app
- **Jenkins**: http://localhost:8080/job/pipeline-from-gitlab-ssh/

---

## 🇬🇧 English Version

### 📖 About the Project

Demonstration project of a **full CI/CD cycle** for a Go application:
from commit to production in Kubernetes. The project shows integration of
**GitLab CI** and **Jenkins**, working with **Docker Hub**, **HashiCorp Vault**
and **Kubernetes** (Minikube).

### 🛠 Tech Stack

| Component | Technology |
|---|---|
| **Language** | Go 1.21 |
| **CI/CD** | GitLab CI, Jenkins (Pipeline as Code) |
| **Containerization** | Docker, Docker Hub |
| **Orchestration** | Kubernetes (Minikube) |
| **Secrets management** | HashiCorp Vault (JWT auth) |
| **Monitoring** | Prometheus, Grafana |
| **GitOps** | Flux |

### 🚀 Pipeline Flow

1. **`git push` to `main`** — GitLab receives the commit.
2. **GitLab CI** runs the pipeline: `build → test → deploy`.
3. **GitLab webhook** sends the event to Jenkins via **localhost.run tunnel**.
4. **Jenkins**:
   - Checkout code
   - Build Docker image in a **Docker agent**
   - Push image to **Docker Hub**
   - Deploy to **Kubernetes**
   - Verify pods and service
5. **Vault** issues secrets to GitLab CI via **JWT authentication**.

### 📁 Project Structure

```
my-pipeline-163/
├── main.go                 # Go application
├── go.mod                  # Go modules
├── Dockerfile              # Multi-stage build
├── Jenkinsfile             # Jenkins Pipeline
├── .gitlab-ci.yml          # GitLab CI config
├── k8s/                    # Kubernetes manifests
│   ├── deployment.yaml
│   └── service.yaml
└── README.md
```

### 🎯 Quick Start

```bash
# 1. Clone the repository
git clone git@gitlab.com:vladimir_vgb/my-pipeline-163.git
cd my-pipeline-163

# 2. Build the Docker image
docker build -t zakon13/my-go-app:latest .

# 3. Run locally
docker run -p 8080:8080 zakon13/my-go-app:latest

# 4. Trigger the pipeline
git push origin main
```

### 🔗 Links

- **GitLab repository**: https://gitlab.com/vladimir_vgb/my-pipeline-163
- **Docker Hub image**: https://hub.docker.com/r/zakon13/my-go-app
- **Jenkins**: http://localhost:8080/job/pipeline-from-gitlab-ssh/

---

## 📄 License / Лицензия

MIT License / Лицензия MIT
