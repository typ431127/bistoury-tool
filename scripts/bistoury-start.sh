#!/bin/sh
# bistoury-agent 启动脚本

function _start() {
    sed -i "s#127.0.0.1:9090#${BISTOURY_SERVICES}#" /bistoury-agent/bin/bistoury-agent-env.sh
    pid=$(ps aux|grep java|grep -v grep|awk '{print $1}')
    /bin/sh /bistoury-agent/bin/bistoury-agent.sh -p ${pid} start
    _register
}
function _register() {
    jobname=$(cat /jenkins_env |grep 'JOB_BASE_NAME'|awk -F '[=]' '{print $2}')
    port=$(cat /jenkins_env |grep 'port'|awk -F '[=]' '{print $2}')
    logdir=$(cat /jenkins_env |grep 'logdir'|awk -F '[=]' '{print $2}')
    gitproject=$(cat jenkins_env |grep giturl|awk -F '[=]' '{print $2}'|sed 's#.*net/##g'|sed 's#.git##g')
    gitbranch=$(cat jenkins_env |grep gitbranch|sed 's#.*/##g')
    mkdir -p /tmp/bistoury/
    echo "project=${gitproject}" > /tmp/bistoury/releaseInfo.properties
    echo "module=." >> /tmp/bistoury/releaseInfo.properties
    echo "output=${gitbranch}" >> /tmp/bistoury/releaseInfo.properties
    curl --location --request POST "${BISTOURY_API}/api/app" \
    --header "Content-Type: application/json" \
    --data-raw "{
        \"code\": \"${jobname}\",
        \"Name\": \"${jobname}\",
        \"Group_code\": \"${jobname}\",
        \"ip\":\"$(hostname -i)\",
        \"port\":${port},
        \"logdir\":\"${logdir}\",
        \"hostname\": \"$(hostname)\"
    }"
}

function _del() {
    curl --location --request DELETE "${BISTOURY_API}/api/app" \
    --header "Content-Type: application/json" \
    --data-raw "{
        \"ip\":\"$(hostname -i)\",
        \"hostname\": \"$(hostname)\"
    }"
}




case $1 in
    start)
        echo "start..."
        _start
        ;;
    del)
        echo "注销实例..."
        _del
        ;;
    register)
        echo "注册实例..."
        _register
        ;;
    *)
        echo "start     启动bistoury-agent"
        echo "del       注销实例"
        echo "register  注册实例"
esac
