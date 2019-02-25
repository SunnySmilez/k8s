#!/bin/sh

project_name=test_project
commit_id=`git rev-parse --short HEAD`
job_name=$project_name"_"$commit_id
#php jenkins.php create $job_name
php jenkins.php build $job_name