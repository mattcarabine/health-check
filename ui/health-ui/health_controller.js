(function () {
    angular.module('health').controller('healthController', healthController);
    healthController.$inject = ['$rootScope', '$scope', '$stateParams', 'healthService', 'mnPoller'];

    function healthController($rootScope, $scope, $stateParams,  healthService, mnPoller) {
        var hc = this;
        hc.openPingDropdowns = {};
        hc.openSummaryDropdowns = [];
        hc.togglePingDropdown = togglePingDropdown;
        hc.toggleSummaryDropdown = toggleSummaryDropdown;
        hc.pingNodeDroppedDown = pingNodeDroppedDown;
        hc.someUnreachable = someUnreachable;

        new mnPoller($scope, function () {
            return healthService.fetchHealth()
        })
            .subscribe(function (health) {
                console.log("Wahoo", health);
                hc.health = health;
            })
            .cycle();

        new mnPoller($scope, function () {
            return healthService.fetchPing()
        })
            .subscribe(function (ping) {
                console.log("Ping:", ping);
                hc.ping = ping.pingReport;
            })
            .cycle();

        function togglePingDropdown(topNode, subNode){
            if (subNode){
                var nodeIndex = hc.openPingDropdowns[topNode].indexOf(subNode);
                if (nodeIndex > -1){
                    hc.openPingDropdowns[topNode].splice(nodeIndex, 1);
                }  else  {
                    hc.openPingDropdowns[topNode].push(subNode);
                }
            } else {
                if (hc.openPingDropdowns[topNode]){
                    delete hc.openPingDropdowns[topNode];
                }  else  {
                    hc.openPingDropdowns[topNode] = [];
                }
            }
        }

        function pingNodeDroppedDown(topNode, subNode){
            return hc.openPingDropdowns[topNode] && (!subNode || hc.openPingDropdowns[topNode].includes(subNode))
        }

        function toggleSummaryDropdown(node){
            var nodeIndex = hc.openSummaryDropdowns.indexOf(node);
            if (nodeIndex > -1){
                hc.openSummaryDropdowns.splice(nodeIndex, 1)
            } else {
                hc.openSummaryDropdowns.push(node)
            }
        }

        function someUnreachable(node){
            return node.some(function(r){return !!r.unreachablePorts});
        }
    }
})();
