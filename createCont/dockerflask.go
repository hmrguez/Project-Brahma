package createCont

var dockerflask = `
# Use the official Python image as the base image
FROM python:3.9-slim-buster

# Set the working directory to /app
WORKDIR /app

# Copy the requirements file and install any dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code
COPY . ./

# Expose port 5000 for the application
EXPOSE 5000

# Start the application
CMD [ "python", "app.py" ]
`
