angular.module('app')
.controller('PackagesCtrl', function ($rootScope, $scope, PackageSvc, CategorySvc) {
  PackageSvc.fetch().success(function (packages) {
    $scope.packages = packages
  })

  CategorySvc.fetch().success(function (categories) {
    $scope.categories = categories
  })
})
