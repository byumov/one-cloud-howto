yum install -y \
    python3
    # nginx
    # httpd
    # ...

# install go
curl -O https://nexus.odkl.ru/repository/dists/docker-images/go1.17.6.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz
rm -rf go1.17.6.linux-amd64.tar.gz

cd /opt/myapp
/usr/local/go/bin/go build app.go
