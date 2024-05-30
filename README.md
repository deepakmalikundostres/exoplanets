# Exoplanets API

## Description
The Exoplanets API allows you to manage a collection of exoplanets. You can add, update, delete, and retrieve information about exoplanets, as well as estimate the fuel needed for a mission based on the exoplanet's characteristics.

## Installation
To get started with this project, follow these steps:

### Using Go

1. Clone the repository:
    ```sh
    git clone https://github.com/deepakmalikundostres/exoplanets.git
    cd exoplanets
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Build the project:
    ```sh
    go build -o exoplanets-api
    ```

4. Run the server:
    ```sh
    ./exoplanets-api
    ```

### Using Docker

1. Clone the repository:
    ```sh
    git clone https://github.com/deepakmalikundostres/exoplanets.git
    cd exoplanets
    ```

2. Build the Docker image:
    ```sh
    docker build -t exoplanets-api .
    ```

3. Run the Docker container:
    ```sh
    docker run -p 8080:8080 exoplanets-api
    ```

## Usage
The API provides the following endpoints:

### Add an Exoplanet
- **URL:** `/exoplanets`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "name": "string",
        "description": "string",
        "distance": 100,
        "radius": 1.0,
        "mass": 1.0,
        "type": "Terrestrial"
    }
    ```
- **Response:**
    ```json
    {
        "id": "1",
        "name": "string",
        "description": "string",
        "distance": 100,
        "radius": 1.0,
        "mass": 1.0,
        "type": "Terrestrial"
    }
    ```

### List All Exoplanets
- **URL:** `/exoplanets`
- **Method:** `GET`
- **Response:**
    ```json
    [
        {
            "id": "1",
            "name": "string",
            "description": "string",
            "distance": 100,
            "radius": 1.0,
            "mass": 1.0,
            "type": "Terrestrial"
        },
        {
            "id": "2",
            "name": "string",
            "description": "string",
            "distance": 200,
            "radius": 2.0,
            "type": "GasGiant"
        }
    ]
    ```

### Get Exoplanet by ID
- **URL:** `/exoplanets/{id}`
- **Method:** `GET`
- **Response:**
    ```json
    {
        "id": "1",
        "name": "string",
        "description": "string",
        "distance": 100,
        "radius": 1.0,
        "mass": 1.0,
        "type": "Terrestrial"
    }
    ```

### Update Exoplanet
- **URL:** `/exoplanets/{id}`
- **Method:** `PUT`
- **Request Body:**
    ```json
    {
        "name": "string",
        "description": "string",
        "distance": 150,
        "radius": 1.1,
        "mass": 1.1,
        "type": "Terrestrial"
    }
    ```
- **Response:**
    ```json
    {
        "id": "1",
        "name": "string",
        "description": "string",
        "distance": 150,
        "radius": 1.1,
        "mass": 1.1,
        "type": "Terrestrial"
    }
    ```

### Delete Exoplanet
- **URL:** `/exoplanets/{id}`
- **Method:** `DELETE`
- **Response:** `204 No Content`

### Estimate Fuel for Mission
- **URL:** `/exoplanets/{id}/fuel?crewCapacity={crewCapacity}`
- **Method:** `GET`
- **Response:**
    ```json
    {
        "fuel": 12345.67
    }
    ```

## License
This project is licensed under the MIT License.
