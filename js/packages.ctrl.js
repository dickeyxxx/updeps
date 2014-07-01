angular.module('app')
.controller('PackagesCtrl', function ($rootScope, $scope, PackageSvc) {
  $rootScope.activeCtrl = 'PackagesCtrl'
  PackageSvc.fetch().success(function (packages) {
    $scope.packages = packages
  })
})
