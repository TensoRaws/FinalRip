#!/bin/bash

# download easytier-core
wget -c https://github.com/EasyTier/EasyTier/releases/download/v1.2.0/easytier-linux-x86_64-v1.2.0.zip
unzip easytier-linux-x86_64-v1.2.0.zip
mv ./easytier-linux-x86_64/easytier-core /usr/local/bin/
rm -rf easytier-linux-x86_64-v1.2.0.zip easytier-linux-x86_64

# get command line arguments from env variables
EASYTIER_COMMAND="${EASYTIER_COMMAND:-d --network-name finalrip --network-secret finalrip -e tcp://easytier.public.kkrainbow.top:11010}"

SERVICE_NAME="easytier"
EXECUTABLE_PATH="/usr/local/bin/easytier-core"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"

# create service file
cat <<EOF | sudo tee $SERVICE_FILE > /dev/null
[Unit]
Description=EasyTier Core Service
After=network.target

[Service]
Type=simple
ExecStart=${EXECUTABLE_PATH} ${EASYTIER_COMMAND}

[Install]
WantedBy=multi-user.target
EOF

sudo chmod +x ${EXECUTABLE_PATH}

sudo systemctl daemon-reload
sudo systemctl enable ${SERVICE_NAME}.service
sudo systemctl start ${SERVICE_NAME}.service

echo "service ${SERVICE_NAME} has been set to start on boot."
