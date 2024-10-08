<h2>Customer Support App using GO Lang</h2>
<p>This is another beneficial project that i have learned many things during development process. In this project i have done some pactices on;</p>
<ul>
    <li>Authentication & Role Based Authorization,</li>
    <li>Json Web Token-(JWT),</li>
    <li>File Things,</li>
    <li>Routing,</li>
    <li>Middlewares,</li>
    <li>SqlServer,</li>
    <li>Project Folder Structure & Organizing Modules,</li>
    <li>Endpoint Model Validations,</li>
</ul>

<p>To test this app you can just clone into your machine and make sure you have installed Go Lang. You need to define some configurations in <b>.env</b> file such as <u>database connection string</u>.Create and .env file first and copy the content of .env-example file init. Than you can apply you settings. Also you need to make sure you have permission to connect database via tcp and SQL Server Authentication because the <b>gorm</b> ORM uses tcp to handle connection  and we have defined user credentials on .env file. If you have done all configurations that you need follow the commands below.</p>

```
cd btelli-customersupport-app
```

```
go run cmd/main.go
```