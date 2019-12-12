# Swagger Petstore
## Version: 1.0.0

**License:** MIT

### /pets

#### GET
##### Summary:

List all pets

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query | How many items to return at one time (default 100) | No | integer |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | An array of pets |

#### POST
##### Summary:

Create a Pet

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Expected response to a valid request |
| 400 | Bad Request |
| 500 | Internal Server Error |

### /pets/{petId}

#### GET
##### Summary:

Info for a specific pet

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| petId | path | The id of the pet to retrieve | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Expected response to a valid request |
| 404 | Not found |

#### DELETE
##### Summary:

Delete a Pet

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| petId | path | The id of the pet to delete | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | ok |
| 404 | Not found |

#### PATCH
##### Summary:

Update a pet

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| petId | path | The id of the pet to update | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Ok |
| 400 | Bad Request |
| 404 | Not found |
| 500 | Internal Server Error |
