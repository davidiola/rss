# rss
Simple golang rss program that aggregates rss endpoint data &amp; sends a text with the message via Twilio
## USAGE
* Make sure to add your own appropriate Twilio Account SID and corresponding Authentication Token in main.go.  You can create a free Twilio account here: [https://www.twilio.com/try-twilio]
## GOALS
* The goal of this project was to demonstrate how easy it is to hit web endpoints and parse the JSON responses in Go
* If you want to receive this text each day without manually running the program I suggest setting up a free Amazon EC2 instance and using linux cron jobs.  Make sure to install the go language on the machine.   
* Open /etc/crontab in vim and add `30 7 * * * ec2-user /home/ec2-user/go_projects/src/rss/main` if you want the program to run at 7:30 am each day, and `5 17 * * * ec2-user /home/ec2-user/go_projects/src/rss/main` for 5:05 pm each day.  Ensure that the time on the linux instance is set correctly for your time zone :)
