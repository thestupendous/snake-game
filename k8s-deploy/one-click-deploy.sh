if kubectl apply -f ns.yaml && kubectl apply -f deploy.yaml ; then
    echo "All resources deployed"
else
    echo "Resource deployment failed"
fi

# kubectl apply -f ns.yaml
# kubectl apply -f deploy.yaml
# if !$? do echo 