angular.module('app')
.controller('AddPackageCtrl', function ($rootScope, $scope, PackageSvc) {
  $rootScope.activeCtrl = 'AddPackageCtrl'

  $scope.pkg = {}
  $scope.create = function (pkg) {
    PackageSvc.create(pkg)
    .success(function () {
      $scope.pkg = {}
    })
  }
})
