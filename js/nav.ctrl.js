angular.module('app')
.controller('NavCtrl', function ($scope, PackageSvc) {
  $scope.refresh = function () {
    PackageSvc.refresh()
  }
})
