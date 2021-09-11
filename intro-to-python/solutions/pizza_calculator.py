def pizza_calculator(people, slices_per_person, slices_per_pizza):
    """ Pizza calculator solution """
    slices_required = people * slices_per_person
    pizzas = slices_required // slices_per_pizza
    
    if slices_required % slices_per_pizza != 0:
        return pizzas + 1
    
    return pizzas