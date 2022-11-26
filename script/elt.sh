#!/bin/sh
# Download db from http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz
# and extract it to the current directory

# Download the file
cd server
echo "Downloading database"
out=$(wget  http://download.srv.cs.cmu.edu/\~enron/enron_mail_20110402.tgz)

# Extract the file
if [ $? -eq 0 ]; then
    echo "Extracting database"
    tar -xzf enron_mail_20110402.tgz
    echo "Database downloaded and extracted"
else
    echo "Error downloading database"
fi
