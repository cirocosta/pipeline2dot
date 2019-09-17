# concourse pipelines to dot graphs

Usage

```bash
fly -t ci get-pipeline -p concourse | \
  pipeline2dot | \
  dot -Tpng > graph.png
```

Result

![](./graph.png)
