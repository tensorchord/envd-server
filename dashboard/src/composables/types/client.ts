/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface AuthPublicKeyAuthRequest {
  /**
   * ClientVersion contains the version string the connecting client sent if any. May be empty if the client did not
   * provide a client version.
   *
   * required: false
   * in: body
   */
  clientVersion?: string;
  /**
   * ConnectionID is an opaque ID to identify the SSH connection in question.
   *
   * required: true
   * in: body
   */
  connectionId?: string;
  /**
   * Environment is a set of key-value pairs provided by the authentication or configuration system and may be
   * exposed by the backend.
   *
   * required: false
   * in: body
   */
  environment?: Record<string, MetadataValue>;
  /**
   * Files is a key-value pair of file names and their content set by the authentication or configuration system
   * and consumed by the backend.
   *
   * required: false
   * in: body
   */
  files?: Record<string, MetadataBinaryValue>;
  /**
   * Metadata is a set of key-value pairs that carry additional information from the authentication and configuration
   * system to the backends. Backends can expose this information as container labels, environment variables, or
   * other places.
   *
   * required: false
   * in: body
   */
  metadata?: Record<string, MetadataValue>;
  /**
   * PublicKey is the key in the authorized key format.
   *
   * required: true
   */
  publicKey?: string;
  /**
   * RemoteAddress is the IP address and port of the user trying to authenticate.
   *
   * required: true
   * in: body
   */
  remoteAddress?: string;
  /**
   * Username is the username provided on login by the client. This may, but must not necessarily match the
   * authenticated username.
   *
   * required: true
   * in: body
   */
  username?: string;
}

export interface AuthResponseBody {
  /**
   * AuthenticatedUsername contains the username that was actually verified. This may differ from LoginUsername when,
   * for example OAuth2 or Kerberos authentication is used. This field is empty until the authentication phase is
   * completed.
   *
   * required: false
   * in: body
   */
  authenticatedUsername?: string;
  /**
   * ClientVersion contains the version string the connecting client sent if any. May be empty if the client did not
   * provide a client version.
   *
   * required: false
   * in: body
   */
  clientVersion?: string;
  /**
   * ConnectionID is an opaque ID to identify the SSH connection in question.
   *
   * required: true
   * in: body
   */
  connectionId?: string;
  /**
   * Environment is a set of key-value pairs provided by the authentication or configuration system and may be
   * exposed by the backend.
   *
   * required: false
   * in: body
   */
  environment?: Record<string, MetadataValue>;
  /**
   * Files is a key-value pair of file names and their content set by the authentication or configuration system
   * and consumed by the backend.
   *
   * required: false
   * in: body
   */
  files?: Record<string, MetadataBinaryValue>;
  /**
   * Metadata is a set of key-value pairs that carry additional information from the authentication and configuration
   * system to the backends. Backends can expose this information as container labels, environment variables, or
   * other places.
   *
   * required: false
   * in: body
   */
  metadata?: Record<string, MetadataValue>;
  /**
   * RemoteAddress is the IP address and port of the user trying to authenticate.
   *
   * required: true
   * in: body
   */
  remoteAddress?: string;
  /**
   * Success indicates if the authentication was successful.
   *
   * required: true
   * in: body
   */
  success?: boolean;
  /**
   * Username is the username provided on login by the client. This may, but must not necessarily match the
   * authenticated username.
   *
   * required: true
   * in: body
   */
  username?: string;
}

export interface ConfigRequest {
  /**
   * AuthenticatedUsername contains the username that was actually verified. This may differ from LoginUsername when,
   * for example OAuth2 or Kerberos authentication is used. This field is empty until the authentication phase is
   * completed.
   *
   * required: false
   * in: body
   */
  authenticatedUsername?: string;
  /**
   * ClientVersion contains the version string the connecting client sent if any. May be empty if the client did not
   * provide a client version.
   *
   * required: false
   * in: body
   */
  clientVersion?: string;
  /**
   * ConnectionID is an opaque ID to identify the SSH connection in question.
   *
   * required: true
   * in: body
   */
  connectionId?: string;
  /**
   * Environment is a set of key-value pairs provided by the authentication or configuration system and may be
   * exposed by the backend.
   *
   * required: false
   * in: body
   */
  environment?: Record<string, MetadataValue>;
  /**
   * Files is a key-value pair of file names and their content set by the authentication or configuration system
   * and consumed by the backend.
   *
   * required: false
   * in: body
   */
  files?: Record<string, MetadataBinaryValue>;
  /**
   * Metadata is a set of key-value pairs that carry additional information from the authentication and configuration
   * system to the backends. Backends can expose this information as container labels, environment variables, or
   * other places.
   *
   * required: false
   * in: body
   */
  metadata?: Record<string, MetadataValue>;
  /**
   * RemoteAddress is the IP address and port of the user trying to authenticate.
   *
   * required: true
   * in: body
   */
  remoteAddress?: string;
  /**
   * Username is the username provided on login by the client. This may, but must not necessarily match the
   * authenticated username.
   *
   * required: true
   * in: body
   */
  username?: string;
}

export interface MetadataBinaryValue {
  /**
   * Sensitive indicates that the metadata value contains sensitive data and should not be transmitted to
   * servers unnecessarily.
   *
   * required: false
   * in: body
   */
  sensitive?: boolean;
  /**
   * Value contains the binary data for the current value.
   *
   * required: true
   * in: body
   * swagger:strfmt: byte
   */
  value?: number[];
}

export interface MetadataValue {
  /**
   * Sensitive indicates that the metadata value contains sensitive data and should not be transmitted to
   * servers unnecessarily.
   */
  sensitive?: boolean;
  /** Value contains the string for the current value. */
  value?: string;
}

export interface TypesAuthNRequest {
  /**
   * LoginName is used to authenticate the user and get
   * an access token for the registry.
   * @example "alice"
   */
  login_name?: string;
  /** Password stores the hashed password. */
  password?: string;
  public_key?: string;
}

export interface TypesAuthNResponse {
  /**
   * An opaque token used to authenticate a user after a successful login
   * Required: true
   * @example "a332139d39b89a241400013700e665a3"
   */
  identity_token?: string;
  /**
   * LoginName is used to authenticate the user and get
   * an access token for the registry.
   * @example "alice"
   */
  login_name?: string;
  /**
   * The status of the authentication
   * Required: true
   * @example "Login successfully"
   */
  status?: string;
}

export interface TypesEnvVar {
  /** Name of the environment variable. Must be a C_IDENTIFIER. */
  name?: string;
  /**
   * Variable references $(VAR_NAME) are expanded
   * using the previously defined environment variables in the container and
   * any service environment variables. If a variable cannot be resolved,
   * the reference in the input string will be unchanged. Double $$ are reduced
   * to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.
   * "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".
   * Escaped references will never be expanded, regardless of whether the variable
   * exists or not.
   * Defaults to "".
   * +optional
   */
  value?: string;
}

export interface TypesEnvironment {
  created_at?: number;
  labels?: Record<string, string>;
  name?: string;
  resource?: TypesResourceSpec;
  spec?: TypesEnvironmentSpec;
  status?: TypesEnvironmentStatus;
}

export interface TypesEnvironmentCreateRequest {
  created_at?: number;
  labels?: Record<string, string>;
  name?: string;
  resource?: TypesResourceSpec;
  spec?: TypesEnvironmentSpec;
  status?: TypesEnvironmentStatus;
}

export interface TypesEnvironmentCreateResponse {
  environment?: TypesEnvironment;
  /**
   * Warnings encountered when creating the pod
   * Required: true
   */
  warnings?: string[];
}

export interface TypesEnvironmentGetResponse {
  created_at?: number;
  labels?: Record<string, string>;
  name?: string;
  resource?: TypesResourceSpec;
  spec?: TypesEnvironmentSpec;
  status?: TypesEnvironmentStatus;
}

export interface TypesEnvironmentListResponse {
  items?: TypesEnvironment[];
}

export interface TypesEnvironmentPort {
  name?: string;
  port?: number;
}

export type TypesEnvironmentRemoveResponse = object;

export interface TypesEnvironmentSpec {
  cmd?: string[];
  env?: TypesEnvVar[];
  image?: string;
  owner?: string;
  ports?: TypesEnvironmentPort[];
}

export interface TypesEnvironmentStatus {
  jupyter_addr?: string;
  phase?: string;
  rstudio_server_addr?: string;
}

export interface TypesImageGetResponse {
  created?: number;
  digest?: string;
  labels?: Record<string, string>;
  /** @example "pytorch-cuda:dev" */
  name?: string;
  size?: number;
}

export interface TypesImageListResponse {
  items?: TypesImageMeta[];
}

export interface TypesImageMeta {
  created?: number;
  digest?: string;
  labels?: Record<string, string>;
  /** @example "pytorch-cuda:dev" */
  name?: string;
  size?: number;
}

export interface TypesResourceSpec {
  cpu?: string;
  gpu?: string;
  memory?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "http://localhost:8080/api/v1";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === "number" ? value : `${value}`)}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key)))
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.Text]: (input: any) => (input !== null && typeof input !== "string" ? JSON.stringify(input) : input),
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
            ? JSON.stringify(property)
            : `${property}`,
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(requestParams.headers || {}),
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : requestParams.signal,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title envd server API
 * @version v0.0.12
 * @license MPL 2.0 (https://mozilla.org/MPL/2.0/)
 * @baseUrl http://localhost:8080/api/v1
 * @contact envd maintainers <envd-maintainers@tensorchord.ai> (https://github.com/tensorchord/envd)
 *
 * envd backend server
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * @description get the status of server.
   *
   * @tags root
   * @name GetRoot
   * @summary Show the status of server.
   * @request GET:/
   */
  getRoot = (params: RequestParams = {}) =>
    this.request<Record<string, any>, any>({
      path: `/`,
      method: "GET",
      format: "json",
      ...params,
    });

  config = {
    /**
     * @description It is called by the containerssh webhook. and is not expected to be used externally.
     *
     * @tags ssh-internal
     * @name ConfigCreate
     * @summary Update the config of containerssh.
     * @request POST:/config
     */
    configCreate: (request: ConfigRequest, params: RequestParams = {}) =>
      this.request<void, any>({
        path: `/config`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        ...params,
      }),
  };
  login = {
    /**
     * @description login to the server.
     *
     * @tags user
     * @name LoginCreate
     * @summary login the user.
     * @request POST:/login
     */
    loginCreate: (request: TypesAuthNRequest, params: RequestParams = {}) =>
      this.request<TypesAuthNResponse, any>({
        path: `/login`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  pubkey = {
    /**
     * @description It is called by the containerssh webhook. and is not expected to be used externally.
     *
     * @tags ssh-internal
     * @name PubkeyCreate
     * @summary authenticate the public key.
     * @request POST:/pubkey
     */
    pubkeyCreate: (request: AuthPublicKeyAuthRequest, params: RequestParams = {}) =>
      this.request<AuthResponseBody, any>({
        path: `/pubkey`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  register = {
    /**
     * @description register the user for the given public key.
     *
     * @tags user
     * @name RegisterCreate
     * @summary register the user.
     * @request POST:/register
     */
    registerCreate: (request: TypesAuthNRequest, params: RequestParams = {}) =>
      this.request<TypesAuthNResponse, any>({
        path: `/register`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  users = {
    /**
     * @description List the environment.
     *
     * @tags environment
     * @name EnvironmentsDetail
     * @summary List the environment.
     * @request GET:/users/{login_name}/environments
     * @secure
     */
    environmentsDetail: (loginName: string, params: RequestParams = {}) =>
      this.request<TypesEnvironmentListResponse, any>({
        path: `/users/${loginName}/environments`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Create the environment.
     *
     * @tags environment
     * @name EnvironmentsCreate
     * @summary Create the environment.
     * @request POST:/users/{login_name}/environments
     * @secure
     */
    environmentsCreate: (loginName: string, request: TypesEnvironmentCreateRequest, params: RequestParams = {}) =>
      this.request<TypesEnvironmentCreateResponse, any>({
        path: `/users/${loginName}/environments`,
        method: "POST",
        body: request,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Remove the environment.
     *
     * @tags environment
     * @name EnvironmentsDelete
     * @summary Remove the environment.
     * @request DELETE:/users/{login_name}/environments/{name}
     * @secure
     */
    environmentsDelete: (loginName: string, name: string, params: RequestParams = {}) =>
      this.request<TypesEnvironmentRemoveResponse, any>({
        path: `/users/${loginName}/environments/${name}`,
        method: "DELETE",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Get the environment with the given environment name.
     *
     * @tags environment
     * @name EnvironmentsDetail2
     * @summary Get the environment.
     * @request GET:/users/{login_name}/environments/{name}
     * @originalName environmentsDetail
     * @duplicate
     * @secure
     */
    environmentsDetail2: (loginName: string, name: string, params: RequestParams = {}) =>
      this.request<TypesEnvironmentGetResponse, any>({
        path: `/users/${loginName}/environments/${name}`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description List the images.
     *
     * @tags image
     * @name ImagesDetail
     * @summary List the images.
     * @request GET:/users/{login_name}/images
     * @secure
     */
    imagesDetail: (loginName: string, name: string, params: RequestParams = {}) =>
      this.request<TypesImageListResponse, any>({
        path: `/users/${loginName}/images`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Get the image with the given image name.
     *
     * @tags image
     * @name ImagesDetail2
     * @summary Get the image.
     * @request GET:/users/{login_name}/images/{name}
     * @originalName imagesDetail
     * @duplicate
     * @secure
     */
    imagesDetail2: (loginName: string, name: string, params: RequestParams = {}) =>
      this.request<TypesImageGetResponse, any>({
        path: `/users/${loginName}/images/${name}`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
