angular.module('app')
.config(function ($routeProvider, $locationProvider) {
  $locationProvider.html5Mode(true)
  $routeProvider
  .when('/', { templateUrl: '/packages.html', controller: 'PackagesCtrl' })
  .when('/package/:path*', { templateUrl: '/package_detail.html', controller: 'PackageDetailCtrl' })
  .when('/admin', { templateUrl: '/admin.html', controller: 'AdminCtrl' })
  .otherwise({redirectTo: '/'})
})
