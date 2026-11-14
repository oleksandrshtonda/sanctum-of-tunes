type HTTPClientParams = {
  baseUrl: string;
}

class HTTPClient {
  private baseUrl: string;

  public constructor(params: HTTPClientParams) {
    this.baseUrl = params.baseUrl;
  }

  public get() {
    const options: RequestInit = {
      method: "GET",
    }

    return fetch(this.baseUrl, options)
  }

  public post() {
    const options: RequestInit = {
      method: "POST",
    }

    return fetch(this.baseUrl, options)
  }

  public put() {
    const options: RequestInit = {
      method: "PUT",
    }

    return fetch(this.baseUrl, options)
  }

  public patch() {
    const options: RequestInit = {
      method: "PATCH",
    }

    return fetch(this.baseUrl, options)
  }

  public delete() {
    const options: RequestInit = {
      method: "DELETE",
    }

    return fetch(this.baseUrl, options)
  }

}

const httpClientParams: HTTPClientParams = {
  baseUrl: "http://localhost:8000/"
}
export const httpClient = new HTTPClient(httpClientParams);
