
## Dockerfile
```
FROM  blang/busybox-bash
ADD /crcservice /
CMD ["/crcservice"]
EXPOSE 8080
```
