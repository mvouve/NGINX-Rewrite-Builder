# NGINX Redirect Writer
This is a quick program written in Go that takes an XLSX input and produces a text file which can be copied within an
existing NGINX configuration block.

I often get spreadsheets detailing redirects to implement on a webpage, I designed this program to make the process go a
little faster.

## Usage
Currently there are a couple global constants which define the names of the row headers that are used they are 
```Current Landing Page``` for the old URL and ```Redirect Landing Page``` for the new URL.

The program itself is invoked with two arguments the XLSX first then the output file name. eg:
```./<bin_name> <xlsx> <output>```


## TODOs
- Make an external configuration file