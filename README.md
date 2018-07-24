# Welcome to Buffalo!

Thank you for choosing Buffalo for your web development needs.

You'll need Docker, Docker Compose and [Buffalo](https://gobuffalo.io) installed.
After you have that, getting started is pretty quick.

First, install Go dependencies with [dep](https://github.com/golang/dep):

```console
dep ensure
```

Then, start up the app:

```console
make dev
buffalo dev
```

To deploy to [ACI](https://azure.microsoft.com/en-us/services/container-instances/):

```console
az container create \
	-n hd \
	-g ${RESOURCE_GROUP} \
	--image arschles/hd \
	--ip-address public \
	--ports 3000 \
	-e "GO_ENV=production" "DATABASE_URL=${DATABASE_URL}"
```

Notes:

- you can substitute `arschles/hd` with your own image name.
- `DATABASE_URL` should be something like this: `postgres://${DB_USERNAME}:${DATABASE_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=require` (`DB_PORT` is usually 5432)

The stuff below is Buffalo's boilerplate readme.

## Database Setup

It looks like you chose to set up your application using a postgres database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start postgres for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a
## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
