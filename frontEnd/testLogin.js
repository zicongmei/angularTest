//if (!("token" in window.sessionStorage) || window.sessionStorage.token.length < 10) {
//    window.location.href = '/login.html';
//}
var app = angular.module('testLogin', []);
app.controller('testLoginCtrl', function($scope, $http) {
    $scope.test = function() {
        $http({
            method: "GET",
            url: "request/test",
            headers: {
                'Authorization': window.sessionStorage.token,
                'Accept': 'application/json',
            }
        }).then(function mySuccess(response) {
            console.log("token correct")
        }, function myError(response) {
            console.log("token incorrect")
            window.location.href = '/login.html';
        });
    }
})