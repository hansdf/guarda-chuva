# guarda-chuva
# Should you bring an umbrella to go out today?

This is my first time working with a proper Go application.
it should be a minimalist CLI tool that displays if you should bring an umbrella to go out(as in, it will rain that day)

It consumes a REST API from openweathermap.org that provides 5 day forecast at any location on the globe.
The great JSON-to-Go tool (https://mholt.github.io/json-to-go/) helps in parsing the info

For now, the coordinates(lat, lon) are fixed for Porto Alegre, Rio Grande do Sul.

In the future I'd like to make a web app version of this(using Gin framework) with a frontend!

![image](https://github.com/user-attachments/assets/15a1d440-b0b0-4245-8bf2-106df927193b)


