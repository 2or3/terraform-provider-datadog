---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_synthetics_test Resource - terraform-provider-datadog"
subcategory: ""
description: |-
  Provides a Datadog synthetics test resource. This can be used to create and manage Datadog synthetics test.
---

# datadog_synthetics_test (Resource)

Provides a Datadog synthetics test resource. This can be used to create and manage Datadog synthetics test.

## Example Usage

```terraform
# Example Usage (Synthetics API test)
# Create a new Datadog Synthetics API/HTTP test on https://www.example.org
resource "datadog_synthetics_test" "test_api" {
  type    = "api"
  subtype = "http"
  request_definition {
    method = "GET"
    url    = "https://www.example.org"
  }
  request_headers = {
    Content-Type   = "application/json"
    Authentication = "Token: 1234566789"
  }
  assertion {
    type     = "statusCode"
    operator = "is"
    target   = "200"
  }
  locations = ["aws:eu-central-1"]
  options_list {
    tick_every = 900

    retry {
      count    = 2
      interval = 300
    }

    monitor_options {
      renotify_interval = 100
    }
  }
  name    = "An API test on example.org"
  message = "Notify @pagerduty"
  tags    = ["foo:bar", "foo", "env:test"]

  status = "live"
}


# Example Usage (Synthetics SSL test)
# Create a new Datadog Synthetics API/SSL test on example.org
resource "datadog_synthetics_test" "test_ssl" {
  type    = "api"
  subtype = "ssl"
  request_definition {
    host = "example.org"
    port = 443
  }
  assertion {
    type     = "certificate"
    operator = "isInMoreThan"
    target   = 30
  }
  locations = ["aws:eu-central-1"]
  options_list {
    tick_every         = 900
    accept_self_signed = true
  }
  name    = "An API test on example.org"
  message = "Notify @pagerduty"
  tags    = ["foo:bar", "foo", "env:test"]

  status = "live"
}


# Example Usage (Synthetics TCP test)
# Create a new Datadog Synthetics API/TCP test on example.org
resource "datadog_synthetics_test" "test_tcp" {
  type    = "api"
  subtype = "tcp"
  request_definition {
    host = "example.org"
    port = 443
  }
  assertion {
    type     = "responseTime"
    operator = "lessThan"
    target   = 2000
  }
  locations = ["aws:eu-central-1"]
  options_list {
    tick_every = 900
  }
  name    = "An API test on example.org"
  message = "Notify @pagerduty"
  tags    = ["foo:bar", "foo", "env:test"]

  status = "live"
}


# Example Usage (Synthetics DNS test)
# Create a new Datadog Synthetics API/DNS test on example.org
resource "datadog_synthetics_test" "test_dns" {
  type    = "api"
  subtype = "dns"
  request_definition {
    host = "example.org"
  }
  assertion {
    type     = "recordSome"
    operator = "is"
    property = "A"
    target   = "0.0.0.0"
  }
  locations = ["aws:eu-central-1"]
  options_list {
    tick_every = 900
  }
  name    = "An API test on example.org"
  message = "Notify @pagerduty"
  tags    = ["foo:bar", "foo", "env:test"]

  status = "live"
}


# Example Usage (Synthetics Browser test)
# Support for Synthetics Browser test steps is limited (see below)
# Create a new Datadog Synthetics Browser test starting on https://www.example.org
resource "datadog_synthetics_test" "test_browser" {
  type = "browser"

  request_definition {
    method = "GET"
    url    = "https://app.datadoghq.com"
  }

  device_ids = ["laptop_large"]
  locations  = ["aws:eu-central-1"]

  options_list {
    tick_every = 3600
  }

  name    = "A Browser test on example.org"
  message = "Notify @qa"
  tags    = []

  status = "paused"

  step {
    name = "Check current url"
    type = "assertCurrentUrl"
    params = jsonencode({
      "check" : "contains",
      "value" : "datadoghq"
    })
  }

  variable {
    type    = "text"
    name    = "MY_PATTERN_VAR"
    pattern = "{{numeric(3)}}"
    example = "597"
  }

  variable {
    type    = "email"
    name    = "MY_EMAIL_VAR"
    pattern = "jd8-afe-ydv.{{ numeric(10) }}@synthetics.dtdg.co"
    example = "jd8-afe-ydv.4546132139@synthetics.dtdg.co"
  }

  variable {
    type = "global"
    name = "MY_GLOBAL_VAR"
    id   = "76636cd1-82e2-4aeb-9cfe-51366a8198a2"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **locations** (Set of String) Array of locations used to run the test. Refer to [Datadog documentation](https://docs.datadoghq.com/synthetics/api_test/#request) for available locations (e.g. `aws:eu-central-1`).
- **name** (String) Name of Datadog synthetics test.
- **status** (String) Define whether you want to start (`live`) or pause (`paused`) a Synthetic test. Valid values are `live`, `paused`.
- **type** (String) Synthetics test type. Valid values are `api`, `browser`.

### Optional

- **api_step** (Block List) Steps for multistep api tests (see [below for nested schema](#nestedblock--api_step))
- **assertion** (Block List) Assertions used for the test. Multiple `assertion` blocks are allowed with the structure below. (see [below for nested schema](#nestedblock--assertion))
- **browser_step** (Block List) Steps for browser tests. (see [below for nested schema](#nestedblock--browser_step))
- **browser_variable** (Block List) Variables used for a browser test steps. Multiple `variable` blocks are allowed with the structure below. (see [below for nested schema](#nestedblock--browser_variable))
- **config_variable** (Block List) Variables used for the test configuration. Multiple `config_variable` blocks are allowed with the structure below. (see [below for nested schema](#nestedblock--config_variable))
- **device_ids** (List of String) Array with the different device IDs used to run the test (only for `browser` tests). Valid values are `laptop_large`, `tablet`, `mobile_small`, `chrome.laptop_large`, `chrome.tablet`, `chrome.mobile_small`, `firefox.laptop_large`, `firefox.tablet`, `firefox.mobile_small`.
- **message** (String) A message to include with notifications for this synthetics test. Email notifications can be sent to specific users by using the same `@username` notation as events.
- **options_list** (Block List, Max: 1) (see [below for nested schema](#nestedblock--options_list))
- **request_basicauth** (Block List, Max: 1) The HTTP basic authentication credentials. Exactly one nested block is allowed with the structure below. (see [below for nested schema](#nestedblock--request_basicauth))
- **request_client_certificate** (Block List, Max: 1) Client certificate to use when performing the test request. Exactly one nested block is allowed with the structure below. (see [below for nested schema](#nestedblock--request_client_certificate))
- **request_definition** (Block List, Max: 1) The synthetics test request. Required if `type = "api"`. (see [below for nested schema](#nestedblock--request_definition))
- **request_headers** (Map of String) Header name and value map.
- **request_query** (Map of String) Query arguments name and value map.
- **set_cookie** (String) Cookies to be used for a browser test request, using the [Set-Cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie) syntax.
- **subtype** (String) The subtype of the Synthetic API test. Defaults to `http`. Valid values are `http`, `ssl`, `tcp`, `dns`, `multi`, `icmp`.
- **tags** (List of String) A list of tags to associate with your synthetics test. This can help you categorize and filter tests in the manage synthetics page of the UI. Default is an empty list (`[]`).

### Read-Only

- **id** (String) The ID of this resource.
- **monitor_id** (Number) ID of the monitor associated with the Datadog synthetics test.

<a id="nestedblock--api_step"></a>
### Nested Schema for `api_step`

Required:

- **name** (String) The name of the step.

Optional:

- **allow_failure** (Boolean) Determines whether or not to continue with test if this step fails.
- **assertion** (Block List) Assertions used for the test. Multiple `assertion` blocks are allowed with the structure below. (see [below for nested schema](#nestedblock--api_step--assertion))
- **extracted_value** (Block List) Values to parse and save as variables from the response. (see [below for nested schema](#nestedblock--api_step--extracted_value))
- **is_critical** (Boolean) Determines whether or not to consider the entire test as failed if this step fails. Can be used only if `allow_failure` is `true`.
- **request_basicauth** (Block List, Max: 1) The HTTP basic authentication credentials. Exactly one nested block is allowed with the structure below. (see [below for nested schema](#nestedblock--api_step--request_basicauth))
- **request_client_certificate** (Block List, Max: 1) Client certificate to use when performing the test request. Exactly one nested block is allowed with the structure below. (see [below for nested schema](#nestedblock--api_step--request_client_certificate))
- **request_definition** (Block List, Max: 1) The request for the api step. (see [below for nested schema](#nestedblock--api_step--request_definition))
- **request_headers** (Map of String) Header name and value map.
- **request_query** (Map of String) Query arguments name and value map.
- **subtype** (String) The subtype of the Synthetic multistep API test step. Valid values are `http`.

<a id="nestedblock--api_step--assertion"></a>
### Nested Schema for `api_step.assertion`

Required:

- **operator** (String) Assertion operator. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)).
- **type** (String) Type of assertion. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)). Valid values are `body`, `header`, `statusCode`, `certificate`, `responseTime`, `property`, `recordEvery`, `recordSome`, `tlsVersion`, `minTlsVersion`, `latency`, `packetLossPercentage`, `packetsReceived`, `networkHop`.

Optional:

- **property** (String) If assertion type is `header`, this is the header name.
- **target** (String) Expected value. Depends on the assertion type, refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test) for details.
- **targetjsonpath** (Block List, Max: 1) Expected structure if `operator` is `validatesJSONPath`. Exactly one nested block is allowed with the structure below. (see [below for nested schema](#nestedblock--api_step--assertion--targetjsonpath))

<a id="nestedblock--api_step--assertion--targetjsonpath"></a>
### Nested Schema for `api_step.assertion.targetjsonpath`

Required:

- **jsonpath** (String) The JSON path to assert.
- **operator** (String) The specific operator to use on the path.
- **targetvalue** (String) Expected matching value.



<a id="nestedblock--api_step--extracted_value"></a>
### Nested Schema for `api_step.extracted_value`

Required:

- **name** (String)
- **parser** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--api_step--extracted_value--parser))
- **type** (String) Property of the Synthetics Test Response to use for the variable. Valid values are `http_body`, `http_header`.

Optional:

- **field** (String) When type is `http_header`, name of the header to use to extract the value.

<a id="nestedblock--api_step--extracted_value--parser"></a>
### Nested Schema for `api_step.extracted_value.parser`

Required:

- **type** (String) Type of parser for a Synthetics global variable from a synthetics test. Valid values are `raw`, `json_path`, `regex`.

Optional:

- **value** (String) Regex or JSON path used for the parser. Not used with type `raw`.



<a id="nestedblock--api_step--request_basicauth"></a>
### Nested Schema for `api_step.request_basicauth`

Required:

- **password** (String, Sensitive) Password for authentication.
- **username** (String) Username for authentication.


<a id="nestedblock--api_step--request_client_certificate"></a>
### Nested Schema for `api_step.request_client_certificate`

Required:

- **cert** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--api_step--request_client_certificate--cert))
- **key** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--api_step--request_client_certificate--key))

<a id="nestedblock--api_step--request_client_certificate--cert"></a>
### Nested Schema for `api_step.request_client_certificate.cert`

Required:

- **content** (String, Sensitive) Content of the certificate.

Optional:

- **filename** (String) File name for the certificate.


<a id="nestedblock--api_step--request_client_certificate--key"></a>
### Nested Schema for `api_step.request_client_certificate.key`

Required:

- **content** (String, Sensitive) Content of the certificate.

Optional:

- **filename** (String) File name for the certificate.



<a id="nestedblock--api_step--request_definition"></a>
### Nested Schema for `api_step.request_definition`

Optional:

- **body** (String) The request body.
- **dns_server** (String) DNS server to use for DNS tests (`subtype = "dns"`).
- **dns_server_port** (Number) DNS server port to use for DNS tests.
- **host** (String) Host name to perform the test with.
- **method** (String) The HTTP method. Valid values are `GET`, `POST`, `PATCH`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`.
- **no_saving_response_body** (Boolean) Determines whether or not to save the response body.
- **number_of_packets** (Number) Number of pings to use per test for ICMP tests (`subtype = "icmp"`) between 0 and 10.
- **port** (Number) Port to use when performing the test.
- **should_track_hops** (Boolean) This will turn on a traceroute probe to discover all gateways along the path to the host destination. For ICMP tests (`subtype = "icmp"`).
- **timeout** (Number) Timeout in seconds for the test. Defaults to `60`.
- **url** (String) The URL to send the request to.



<a id="nestedblock--assertion"></a>
### Nested Schema for `assertion`

Required:

- **operator** (String) Assertion operator. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)).
- **type** (String) Type of assertion. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)). Valid values are `body`, `header`, `statusCode`, `certificate`, `responseTime`, `property`, `recordEvery`, `recordSome`, `tlsVersion`, `minTlsVersion`, `latency`, `packetLossPercentage`, `packetsReceived`, `networkHop`.

Optional:

- **property** (String) If assertion type is `header`, this is the header name.
- **target** (String) Expected value. Depends on the assertion type, refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test) for details.
- **targetjsonpath** (Block List, Max: 1) Expected structure if `operator` is `validatesJSONPath`. Exactly one nested block is allowed with the structure below. (see [below for nested schema](#nestedblock--assertion--targetjsonpath))

<a id="nestedblock--assertion--targetjsonpath"></a>
### Nested Schema for `assertion.targetjsonpath`

Required:

- **jsonpath** (String) The JSON path to assert.
- **operator** (String) The specific operator to use on the path.
- **targetvalue** (String) Expected matching value.



<a id="nestedblock--browser_step"></a>
### Nested Schema for `browser_step`

Required:

- **name** (String) Name of the step.
- **params** (Block List, Min: 1, Max: 1) Parameters for the step. (see [below for nested schema](#nestedblock--browser_step--params))
- **type** (String) Type of the step. Valid values are `assertCurrentUrl`, `assertElementAttribute`, `assertElementContent`, `assertElementPresent`, `assertEmail`, `assertFileDownload`, `assertFromJavascript`, `assertPageContains`, `assertPageLacks`, `click`, `extractFromJavascript`, `extractVariable`, `goToEmailLink`, `goToUrl`, `goToUrlAndMeasureTti`, `hover`, `playSubTest`, `pressKey`, `refresh`, `runApiTest`, `scroll`, `selectOption`, `typeText`, `uploadFiles`, `wait`.

Optional:

- **allow_failure** (Boolean) Determines if the step should be allowed to fail.
- **force_element_update** (Boolean) Force update of the "element" parameter for the step
- **timeout** (Number) Used to override the default timeout of a step.

<a id="nestedblock--browser_step--params"></a>
### Nested Schema for `browser_step.params`

Optional:

- **attribute** (String) Name of the attribute to use for an "assert attribute" step.
- **check** (String) Check type to use for an assertion step. Valid values are `equals`, `notEquals`, `contains`, `notContains`, `startsWith`, `notStartsWith`, `greater`, `lower`, `greaterEquals`, `lowerEquals`, `matchRegex`, `between`, `isEmpty`, `notIsEmpty`.
- **click_type** (String) Type of click to use for a "click" step.
- **code** (String) Javascript code to use for the step.
- **delay** (Number) Delay between each key stroke for a "type test" step.
- **element** (String) Element to use for the step, json encoded string.
- **email** (String) Details of the email for an "assert email" step.
- **file** (String) For an "assert download" step.
- **files** (String) Details of the files for an "upload files" step, json encoded string.
- **modifiers** (List of String) Modifier to use for a "press key" step.
- **playing_tab_id** (String) ID of the tab to play the subtest.
- **request** (String) Request for an API step.
- **subtest_public_id** (String) ID of the Synthetics test to use as subtest.
- **value** (String) Value of the step.
- **variable** (Block List, Max: 1) Details of the variable to extract. (see [below for nested schema](#nestedblock--browser_step--params--variable))
- **with_click** (Boolean) For "file upload" steps.
- **x** (Number) X coordinates for a "scroll step".
- **y** (Number) Y coordinates for a "scroll step".

<a id="nestedblock--browser_step--params--variable"></a>
### Nested Schema for `browser_step.params.variable`

Optional:

- **example** (String) Example of the extracted variable.
- **name** (String) Name of the extracted variable.




<a id="nestedblock--browser_variable"></a>
### Nested Schema for `browser_variable`

Required:

- **name** (String) Name of the variable.
- **type** (String) Type of browser test variable. Valid values are `element`, `email`, `global`, `javascript`, `text`.

Optional:

- **example** (String) Example for the variable.
- **id** (String) ID of the global variable to use. This is actually only used (and required) in the case of using a variable of type `global`.
- **pattern** (String) Pattern of the variable.


<a id="nestedblock--config_variable"></a>
### Nested Schema for `config_variable`

Required:

- **name** (String) Name of the variable.
- **type** (String) Type of test configuration variable. Valid values are `global`, `text`.

Optional:

- **example** (String) Example for the variable.
- **id** (String) When type = `global`, ID of the global variable to use.
- **pattern** (String) Pattern of the variable.


<a id="nestedblock--options_list"></a>
### Nested Schema for `options_list`

Required:

- **tick_every** (Number) How often the test should run (in seconds). Valid values are `30`, `60`, `300`, `900`, `1800`, `3600`, `21600`, `43200`, `86400`, `604800`.

Optional:

- **accept_self_signed** (Boolean) For SSL test, whether or not the test should allow self signed certificates.
- **allow_insecure** (Boolean) Allows loading insecure content for an HTTP test.
- **follow_redirects** (Boolean) For API HTTP test, whether or not the test should follow redirects.
- **min_failure_duration** (Number) Minimum amount of time in failure required to trigger an alert. Default is `0`.
- **min_location_failed** (Number) Minimum number of locations in failure required to trigger an alert. Default is `1`.
- **monitor_name** (String) The monitor name is used for the alert title as well as for all monitor dashboard widgets and SLOs.
- **monitor_options** (Block List, Max: 1) (see [below for nested schema](#nestedblock--options_list--monitor_options))
- **monitor_priority** (Number)
- **no_screenshot** (Boolean) Prevents saving screenshots of the steps.
- **retry** (Block List, Max: 1) (see [below for nested schema](#nestedblock--options_list--retry))

<a id="nestedblock--options_list--monitor_options"></a>
### Nested Schema for `options_list.monitor_options`

Optional:

- **renotify_interval** (Number) Specify a renotification frequency.


<a id="nestedblock--options_list--retry"></a>
### Nested Schema for `options_list.retry`

Optional:

- **count** (Number) Number of retries needed to consider a location as failed before sending a notification alert.
- **interval** (Number) Interval between a failed test and the next retry in milliseconds.



<a id="nestedblock--request_basicauth"></a>
### Nested Schema for `request_basicauth`

Required:

- **password** (String, Sensitive) Password for authentication.
- **username** (String) Username for authentication.


<a id="nestedblock--request_client_certificate"></a>
### Nested Schema for `request_client_certificate`

Required:

- **cert** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--request_client_certificate--cert))
- **key** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--request_client_certificate--key))

<a id="nestedblock--request_client_certificate--cert"></a>
### Nested Schema for `request_client_certificate.cert`

Required:

- **content** (String, Sensitive) Content of the certificate.

Optional:

- **filename** (String) File name for the certificate.


<a id="nestedblock--request_client_certificate--key"></a>
### Nested Schema for `request_client_certificate.key`

Required:

- **content** (String, Sensitive) Content of the certificate.

Optional:

- **filename** (String) File name for the certificate.



<a id="nestedblock--request_definition"></a>
### Nested Schema for `request_definition`

Optional:

- **body** (String) The request body.
- **dns_server** (String) DNS server to use for DNS tests (`subtype = "dns"`).
- **dns_server_port** (Number) DNS server port to use for DNS tests.
- **host** (String) Host name to perform the test with.
- **method** (String) The HTTP method. Valid values are `GET`, `POST`, `PATCH`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`.
- **no_saving_response_body** (Boolean) Determines whether or not to save the response body.
- **number_of_packets** (Number) Number of pings to use per test for ICMP tests (`subtype = "icmp"`) between 0 and 10.
- **port** (Number) Port to use when performing the test.
- **should_track_hops** (Boolean) This will turn on a traceroute probe to discover all gateways along the path to the host destination. For ICMP tests (`subtype = "icmp"`).
- **timeout** (Number) Timeout in seconds for the test. Defaults to `60`.
- **url** (String) The URL to send the request to.

## Import

Import is supported using the following syntax:

```shell
# Synthetics tests can be imported using their public string ID, e.g.
terraform import datadog_synthetics_test.fizz abc-123-xyz
```
