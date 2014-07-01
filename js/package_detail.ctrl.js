angular.module('app')
.controller('PackageDetailCtrl', function ($scope, $rootScope, $location, PackageSvc) {
  $rootScope.activeCtrl = 'PackageDetailCtrl'

  var path = $location.path().split('/').splice(2).join('/')
  PackageSvc.find(path).success(function (pkg) {
    $scope.pkg = pkg
  })
})
