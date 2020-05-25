(function() {

    angular.module('health').factory('healthService', getHealthService);

    getHealthService.$inject = ['$http'];

    function getHealthService($http) {

        return {
            fetchHealth: fetchHealth,
            fetchPing: fetchPing
        };

        function fetchHealth(){
            return $http({
                method: 'GET',
                url: '/_p/health/summary'
            }).then(function (resp) {
                return resp.data;
            });
        }

        function fetchPing(){
            return $http({
                method: 'GET',
                url: '/_p/health/ping'
            }).then(function (resp) {
                return resp.data;
            });
        }
    }


})();