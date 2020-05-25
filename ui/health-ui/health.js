(function () {
    "use strict";
    angular
    .module("health",
        ["ui.router", "mnPluggableUiRegistry"])
    .config(function($stateProvider, mnPluggableUiRegistryProvider) {
        $stateProvider
            .state('app.admin.health', {
                abstract: true,
                url: '/health',
                views: {
                    "main@app.admin": {
                        controller: 'healthController as hc',
                        templateUrl:  '/_p/ui/health/health.html'
                    }
                },
                data: {
                    title: "Health Check"
                }
            });
        $stateProvider
            .state('app.admin.health.summary', {
                url: '/summary',
                controller: 'healthController as hc',
                templateUrl: '/_p/ui/health/summary.html',
                data: {
                    title: "Health"
                }
            })
            .state('app.admin.health.connectivity', {
                url: '/connectivity',
                templateUrl: '/_p/ui/health/connectivity.html',
                controller: 'healthController as hc',
                data: {
                    title: "Health"
                }
            });
        mnPluggableUiRegistryProvider.registerConfig({
            name: 'Health',
            state: 'app.admin.health.summary',
            includedByState: 'app.admin.health',
            plugIn: 'workbenchTab',
            index: 2,
            ngShow: "rbac.cluster.bucket['.'].n1ql.index.read"
        });
    });

    angular.module('mnAdmin').requires.push('health');
}());