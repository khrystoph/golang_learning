# S3 URL Signer

## Overview
This tool generates a download URL for an object in a user-defined s3 bucket.
Currently, this requires that you know the object's key as well as the bucket. 
Future versions of this tool will get all objects in a bucket and allow you to select which objects you want a URL for.
Also, this tool will be paired with a tool to upload an image and this will tie into an API as well as hook into a web frontend
that will both allow uploading and downloading of files as long as you are an authorized user.

### Version One
Version one is the generation of URLs that are signed given definted inputs. It takes only a single bucket and single s3 object key. 
This can be iterated on by a script or another tool that will call the function as defined in this package.

### Version Two
Version two will implement a bucket read, list all object keys within the bucket, and allow the user to select the key(s) they wish 
to download. Then, it will allow you to select if you wish to download the object or not.

### Versions Three and Beyond
Not sure what all is planned beyond that point. As I come up with additional functionality I wish to add, I will update version three 
and subsequent versions as I continue to add functionality.

### Usage
```bash
$ go build s3urlsigner.go
$ ./s3urlsigner -b <some bucket name> -k <some s3 object key> -t <time in minutes for url to be valid. Defaults to 1 week.>
```
