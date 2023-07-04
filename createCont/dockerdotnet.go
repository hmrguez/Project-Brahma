package createCont

var dockerdotnet = `
# Use the official .NET Core SDK image as the base image
FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build-env
WORKDIR /app

# Copy the .csproj file and restore any dependencies
COPY *.csproj ./
RUN dotnet restore

# Copy the rest of the application code
COPY . ./

# Build the application in release mode and publish it to /app/publish
RUN dotnet publish -c Release -o /app/publish

# Use the official .NET Core runtime image as the base image
FROM mcr.microsoft.com/dotnet/aspnet:5.0

# Set the working directory to /app
WORKDIR /app
COPY --from=build-env /app/publish .

# Expose port 80 for the application
EXPOSE 80

# Start the application
ENTRYPOINT ["dotnet", "MyApp.dll"]
`