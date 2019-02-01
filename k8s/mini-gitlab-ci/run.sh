kubectl apply -f namespace.yml
# https://docs.gitlab.com/ee/install/kubernetes/gitlab_chart.html
# https://charts.gitlab.io/
helm repo add gitlab https://charts.gitlab.io/
helm repo update
helm upgrade --install gitlab gitlab/gitlab \
  --timeout 600 \
  --set gitlab-runner.runners.namespace=gitlab \
  --set global.hosts.domain=minikube.local \
  --set global.hosts.externalIP=10.10.10.10 \
  --set certmanager-issuer.email=exflyg@gmail.com \
  --set gitlab.migrations.image.repository=registry.gitlab.com/gitlab-org/build/cng/gitlab-rails-ce \
  --set gitlab.sidekiq.image.repository=registry.gitlab.com/gitlab-org/build/cng/gitlab-sidekiq-ce \
  --set gitlab.unicorn.image.repository=registry.gitlab.com/gitlab-org/build/cng/gitlab-unicorn-ce \
  --set gitlab.unicorn.workhorse.image=registry.gitlab.com/gitlab-org/build/cng/gitlab-workhorse-ce \
  --set gitlab.task-runner.image.repository=registry.gitlab.com/gitlab-org/build/cng/gitlab-task-runner-ce \
  --debug