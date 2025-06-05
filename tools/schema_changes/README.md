This utility is designed to quickly compare changes that have occurred in ROS of specific versions.
The output can be used for information or to create an issue.

ROS schema files are required for the utility to work. You can download the latest versions with the following command:

```
export v=7.19
curl "https://tikoci.github.io/restraml/$v/inspect.json" -o- | gzip > ros-$v.json.gz
```

Startup:

```
cd tools/schema_changes
go run . -r 7.18:7.19 -markdown
```

    * ```-r``` - ROS versions to be compared (```-r 7.18:7.19```)
    * ````-f``` - resource path filter (````-f bgp```)
    * ```-all``` - output all changed resources, not just those present in the provider
    * ```-markdown``` - output in Markdown format for copying to GitHub
