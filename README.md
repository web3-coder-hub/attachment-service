**Attachment Microservice API Repository**

This microservice is designed to handle attachments, providing functionality to convert HTML to PDF and generate CSV files. It offers a robust and efficient solution for attachment management in your applications.

### Features:

- **HTML to PDF Conversion:** Convert HTML documents into PDF format seamlessly. This feature is particularly useful for generating printable versions of web content or reports.

- **CSV Generation:** Generate CSV files from data sources, facilitating easy data interchange and compatibility with various applications.

### Getting Started:

To begin using the Attachment Microservice API, follow these simple steps:

1. **Installation:** Clone this repository to your local environment or include it as a dependency in your project.

2. **Configuration:** Configure the microservice according to your requirements, including specifying settings for PDF generation and CSV formatting.

3. **Integration:** Integrate the API into your application by making requests to the designated endpoints for HTML to PDF conversion and CSV generation.

### API Endpoints:

- `/api/v1/createPDF`: Endpoint for converting HTML documents to PDF format.
  - **Method:** POST
  - **Request Body:** HTML content
  - **Response:** PDF file

- `/api/v1/createCSV`: Endpoint for generating CSV files.
  - **Method:** POST
  - **Request Body:** Data for CSV generation
  - **Response:** CSV file

### Usage Example:

```bash
# Convert HTML to PDF
curl -X POST -H "Content-Type: text/html" --data-binary "@example.html" https://your-api-url/api/v1/createPDF > example.pdf

# Generate CSV
curl -X POST -H "Content-Type: application/json" --data '{"data": [...]}' https://your-api-url/api/v1/createCSV > example.csv
```


### Contribution:

Contributions to this repository are welcomed! Feel free to open issues for bug fixes, feature requests, or submit pull requests to enhance the functionality of the microservice.

### License:

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Thank you for using the Attachment Microservice API! If you have any questions or need assistance, please don't hesitate to reach out.