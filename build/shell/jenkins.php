<?php
/**
 * 封装的关于jenkins的操作
 */

/**
 * 获取crumb
 * @param $baseUrl
 * @return mixed
 */
function enableCrumbs($baseUrl) {
    $crumbResult = requestCrumb($baseUrl);
    $data['crumb']             = $crumbResult->crumb;
    $data['crumbRequestField'] = $crumbResult->crumbRequestField;
    return $data;
}

function requestCrumb($baseUrl) {
    $url = sprintf('%s/crumbIssuer/api/json', $baseUrl);

    $curl = curl_init($url);

    curl_setopt($curl, \CURLOPT_RETURNTRANSFER, 1);

    $ret = curl_exec($curl);

    $crumbResult = json_decode($ret);

    if (!$crumbResult instanceof \stdClass) {
        throw new \RuntimeException('Error during json_decode of csrf crumb');
    }

    return $crumbResult;
}

/**
 * 创建job
 * @param $baseUrl
 * @param $jobname
 * @param $xmlConfiguration
 * @param $crumbRequestField
 * @param $crumb
 * @return bool
 * @throws Exception
 */
function createJob($baseUrl, $jobname, $xmlConfiguration, $crumbRequestField, $crumb) {
    $url  = sprintf('%s/createItem?name=%s', $baseUrl, $jobname);
    $curl = curl_init($url);
    curl_setopt($curl, \CURLOPT_POST, 1);

    curl_setopt($curl, \CURLOPT_POSTFIELDS, $xmlConfiguration);
    curl_setopt($curl, \CURLOPT_RETURNTRANSFER, 1);

    $headers = array('Content-Type: text/xml');

    $headers[] = "$crumbRequestField: $crumb";

    curl_setopt($curl, \CURLOPT_HTTPHEADER, $headers);

    $response = curl_exec($curl);

    if (curl_getinfo($curl, CURLINFO_HTTP_CODE) != 200) {
        throw new \Exception(sprintf('Job %s already exists', $jobname));
    }

    if (curl_errno($curl)) {
        throw new \Exception(sprintf('Error creating job %s', $jobname));
    }

    return true;
}

/**
 * build
 * @param $baseUrl
 * @param $jobName
 * @param $crumbRequestField
 * @param $crumb
 * @param array $parameters
 * @return bool
 */
function launchJob($baseUrl, $jobName, $crumbRequestField, $crumb, $parameters = array()) {
    if (0 === count($parameters)) {
        $url = sprintf('%s/job/%s/build', $baseUrl, $jobName);
    } else {
        $url = sprintf('%s/job/%s/buildWithParameters', $baseUrl, $jobName);
    }

    $curl = curl_init($url);

    curl_setopt($curl, \CURLOPT_POST, 1);
    curl_setopt($curl, \CURLOPT_POSTFIELDS, http_build_query($parameters));

    $headers = array();

    $headers[] = "$crumbRequestField: $crumb";

    curl_setopt($curl, \CURLOPT_HTTPHEADER, $headers);

    curl_exec($curl);

    return true;
}

$jobName = $argv[2];
if ($argv[1] == 'create') {
    $github = "https://github.com/SunnySmilez/k8s";

    $build_sh = file_get_contents("build.sh");
    $build_sh = str_replace("{jobName}", $jobName, $build_sh);
    $jenkins_xml = file_get_contents("jenkins.xml");
    $jenkins_xml = str_replace("{github}", $github, $jenkins_xml);
    $jenkins_xml = str_replace("{build_shell}", $build_sh, $jenkins_xml);
    $data = enableCrumbs("http://zhouzhi:123123@127.0.0.1:9090");
    return createJob("http://zhouzhi:123123@127.0.0.1:9090", $jobName, $jenkins_xml, $data['crumbRequestField'], $data['crumb']);
} elseif($argv[1] == 'build') {
    $data = enableCrumbs("http://zhouzhi:123123@127.0.0.1:9090");
    return launchJob("http://zhouzhi:123123@127.0.0.1:9090", $jobName, $data['crumbRequestField'], $data['crumb']);
} else {
    return "usage:\r\ncreate|build jonName";
}