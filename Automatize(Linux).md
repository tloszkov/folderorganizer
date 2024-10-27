To run this script every evening at 10 PM on Arch Linux, you can use a systemd timer or a cron job. Here’s how to set it up with both options:
# Option 1: Using a Systemd Timer

## 1. Create a Service File

Create a service file for your script in ``/etc/systemd/system/organize-folder.service:``
```bash
sudo nano /etc/systemd/system/organize-folder.service
```

Add the following content to define the service. Be sure to replace /path/to/your/script with the actual path to your Go program:

```
[Unit]
Description=Organize Folder at 10 PM

[Service]
ExecStart=/usr/bin/go run /path/to/your/script/main.go
```

## 2. Create a Timer File

Now, create a timer file in ``/etc/systemd/system/organize-folder.timer:``
```bash
sudo nano /etc/systemd/system/organize-folder.timer
```

Add the following content to schedule the task to run daily at 10 PM:

```
[Unit]
Description=Run Organize Folder Script Daily at 10 PM

[Timer]
OnCalendar=*-*-* 22:00:00
Persistent=true

[Install]
WantedBy=timers.target
```

## 3. Enable and Start the Timer
To enable and start the timer:
```bash
sudo systemctl enable organize-folder.timer
sudo systemctl start organize-folder.timer
```

## 4. Check the Timer Status

You can check if the timer is active and scheduled to run by using:
```bash
systemctl list-timers --all
```

# Easier way: Using Cron

Edit the Cron Jobs
Open the crontab editor:
```bash
crontab -e
```

Add the Cron Job
Add this line to schedule the script to run daily at 10 PM. Make sure to replace /path/to/your/script/main.go with the path to your Go program:
```
0 22 * * * /usr/bin/go run /path/to/your/script/main.go
```

Save and Exit

This job will now run every day at 10 PM.