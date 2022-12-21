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
  annotations?: Record<string, string>;
  created_at?: number;
  labels?: Record<string, string>;
  name?: string;
  resource?: TypesResourceSpec;
  spec?: TypesEnvironmentSpec;
  status?: TypesEnvironmentStatus;
}

export interface TypesEnvironmentCreateRequest {
  annotations?: Record<string, string>;
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
  annotations?: Record<string, string>;
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
  apt_packages?: string[];
  cmd?: string[];
  env?: TypesEnvVar[];
  image?: string;
  owner?: string;
  ports?: TypesEnvironmentPort[];
  pypi_commands?: string[];
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
