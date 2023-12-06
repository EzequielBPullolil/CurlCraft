# CurlCraft
A Go-developed CLI aimed at enhancing the CURL tool by adding functionalities to improve the user experience and overall usability.

## Use
Most operations that can be carried out in CURL can be performed in CurlCraft with minimal adjustments.

However, there is a notable divergence in behavior. Unlike CURL, CurlCraft does not automatically produce an HTML response unless explicitly indicated. By default, it furnishes only the header response. The default output from the request encompasses the following data:

    Status: <http_status>
    HTTP: HTTP/HTTPs
    Content-type: Json/html/etc
    x-request-id: 
    set-cookie: {
      "cookie1": <value>,
      "cookie2": <value>
    }

## Basic Request
In CurlCraft, akin to CURL, when a request is made to an HTTP/HTTPS link without specifying parameters, it defaults to a GET request. 
```Bash
  CurlCraft http://ejemplo.com
```
To obtain the file/response from the link, you need to make the request by adding the -h flag.
```Bash
  CurlCraft http://ejemplo.com -b 
```
## Request with http methods
In CurlCraft, requests can be created with any HTTP method using its keyword.

#### Request
```Bash
  CurlCraft http://ejemplo.com [METHOD]
```
The method can be any of those used in HTTP.

### Request with data
Just like in Curl, you can send data in requests with formats like form, JSON, XML, etc, in CurlCraft. After specifying the link, you need to indicate the desired format and use JSON syntax to specify the key-value pair of the data you want the request to include.
```Bash
  CurlCraft http://ejemplo.com [METHOD] [content-type] -d {
    "key": value,
    "key2": value
  }
```

### Comparison of a complete request between Curl and Curl Craft
```Bash
curl http://ejemplo.com -X POST -d "name=valor_aleatorio&password=valor_aleatorio" 
```

```Bash
CurlCraft http://ejemplo.com POST  FORM -d  '{"name": valor_aleatorio, "passowrd": valor_aleatorio }'
```
Please be aware that, even if the content-type is set to FORM, the body syntax consistently follows the JSON format in all requests. Why is this the case? It is designed to alleviate users from the burden of implementing intricate syntax when creating requests.

This becomes even more evident when the request employs a content-type of XML.
```Bash
curl -X POST -H "Content-Type: application/xml" -d '<data><name>valor_aleatorio</name><password>valor_aleatorio</password></data>' http://ejemplo.com

```
```Bash
CurlCraft POST http://ejemplo.com FORM -d  '{ "name": valor_aleatorio,"password": valor_aleatorio }'
```
## Complex Request
When dealing with making requests to a link, one often needs to send multiple requests to the same link with different HTTP methods. Well, this can be addressed using the -complex flag.


```Bash
CurlCraft http://ejemplo.com -complex POST GET PUT DELETE FORM -d  '{"name": valor_aleatorio,"password": valor_aleatorio }'
```

By utilizing the -complex flag, you have the flexibility to supply the command with multiple distinct HTTP methods following the -complex flag. Notably, the order holds significance. If you furnish the command with methods such as POST, GET, PUT, DELETE, it will adhere to the specified sequence—initiating a POST request initially, progressing to the subsequent GET request after the completion of each preceding request, and continuing accordingly.

As you can see, the request can also include sending values and still function seamlessly.
