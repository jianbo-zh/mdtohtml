---
Tags: kubectl, auto-completion
---
# kubectl自动补全（Ubuntu）
```
kubectl completion bash > ~/.kube/completion.bash.inc
printf "
# Kubectl shell completion
source '$HOME/.kube/completion.bash.inc'
" >> $HOME/.bashrc
source $HOME/.bashrc
```