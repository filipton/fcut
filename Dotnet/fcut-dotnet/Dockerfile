FROM mcr.microsoft.com/dotnet/aspnet:7.0
WORKDIR /app
COPY ./bin/Release/net7.0/publish/ /app/

ENV ASPNETCORE_URLS=http://0.0.0.0:8001
EXPOSE 8001

CMD ["dotnet", "fcut-dotnet.dll"]
