class Route:
    def __init__(self, start_input, end_input, num_input):
        """ (Route, str, str, float) -> None
        Initialize self variables with the given inputs
        """
        self.start = str(start_input)
        self.end = str(end_input)
        self.num = int(num_input)

    def __repr__(self):
        """(Route) -> str
        Return canonical representation of Route
        """
        return "Route('" + self.start + "', '" + self.end + "', " + str(self.num) + ")"


    def modify_num_daily_flights_by(self, num_input):
        """ (Route, int) -> None
        Initialize self variables with the given inputs
        """
        self.num += int(num_input)


class Airline:

    def __init__(self, name_input, list_of_routes_input = []):
        """ (Airline, str, list of routes) -> None
        Initialize self variables with the given inputs
        """
        self.name = str(name_input)
        self.routes = list_of_routes_input

    def multiple_flights(self):
        """ (Airline) -> None
        return list with routes that have more than one flight
        """
        m = []
        for r in self.routes:
            if r.num > 1:
                m.append(r)
        return m