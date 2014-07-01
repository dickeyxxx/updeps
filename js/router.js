angular.module('app')
.config(function ($routeProvider, $locationProvider) {
  $locationProvider.html5Mode(true)
  $routeProvider
  .when('/', { templateUrl: '/packages.html', controller: 'PackagesCtrl' })
  .when('/add', { templateUrl: '/add_package.html', controller: 'AddPackageCtrl' })
  .when('/package/:path*', { templateUrl: '/package_detail.html', controller: 'PackageDetailCtrl' })
  .otherwise({redirectTo: '/'})
})
