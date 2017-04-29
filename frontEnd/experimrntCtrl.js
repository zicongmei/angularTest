app.controller('experimrntCtrl', function($scope, $location, $http) {
    $scope.click = [];
    $scope.firstPlayer = false;
    $scope.mouseLocation = function(myE) {
        $scope.x = myE.clientX;
        $scope.y = myE.clientY;
    }
    $scope.mouseClick = function(row, col) {
        var loc = {
            x: row,
            y: col,
            firstPlayer: $scope.firstPlayer
        };
        $scope.click.push(loc)
        $scope.firstPlayer = !$scope.firstPlayer
        $http({
            method: "PUT",
            url: "request/move?row=" + row + "&col=" + col
        }).then(function mySucces(response) {
            $scope.myWelcome = response.data;
        }, function myError(response) {
            $scope.myWelcome = response.statusText;
        });
    }
});