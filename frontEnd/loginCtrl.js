app.controller('loginCtrl', function($scope, $http) {
    $scope.user = "unknown";
    $scope.guestLogin = function(name) {
        console.log("login:" + name)
        $scope.user = name;
        $http({
            method: "POST",
            url: "authenticate/login",
            data: {
                "user": name
            },
        }).then(function mySucces(response) {
            window.sessionStorage.token = response.data
            window.location.href = '/experiment.html';
        }, function myError(response) {
            console.log("Failed to login $http.$$url")
        });
    }
});