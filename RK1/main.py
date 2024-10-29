from typing import Self, Any, Callable


class Conductor:
    def __init__(self, conductorID, name, work_experience, orchestraID) -> Self:
        self._conductorID = conductorID
        self._name = name
        self._work_experience = work_experience
        self._orchestraID = orchestraID

    def __str__(self) -> str:
        return f"Conductor: {self._name} from orchestra\n"

    @property
    def get_conductorID(self): return self._conductorID

    @property
    def get_name(self): return self._name

    @property
    def get_work_experience(self): return self._work_experience


class Orchestra:
    def __init__(self, orchestraID, name, date):
        self.ID = orchestraID
        self._name = name
        self._establishment = date

    def __str__(self):
        return f"Orchestra: {self._name}\n"

    @property
    def get_name(self): return self._name

    @property
    def get_establishment(self): return self._establishment


class Ensemble:
    def __init__(self, conductorID, orchestraID):
        self._conductor_ID = conductorID
        self._orchestra_ID = orchestraID


orchestras = [
    Orchestra(1, "Gewandhausorchester Leipzig", 1781),
    Orchestra(2, "New York Philharmonic", 1842),
    Orchestra(3, "Munich Philharmonic", 1893),
    Orchestra(4, "Boston Symphony Orchestra", 1881),
    Orchestra(5, "London Symphony Orchestra", 1904)
]

conductors = [
    Conductor(1, "Mario", 11, 1),
    Conductor(2, "Zuck", 7, 2),
    Conductor(3, "Enzo", 17, 3),
    Conductor(4, "Edward", 9, 4),
    Conductor(5, "Joshua", 8, 5)
]

ensembles = [
    Ensemble(1, 4),
    Ensemble(2, 1),
    Ensemble(3, 4),
    Ensemble(5, 2),
    Ensemble(2, 2),
    Ensemble(1, 3),
    Ensemble(1, 5),
    Ensemble(4, 4),
    Ensemble(5, 4)
]

one_to_many = [(conductor._name, conductor._work_experience, orchestra._name, orchestra._establishment)
               for conductor in conductors
               for orchestra in orchestras
               if conductor._orchestraID == orchestra.ID]

many_to_many_temp = [(orchestra._name, ensemble._conductor_ID, ensemble._orchestra_ID)
                for orchestra in orchestras
                for ensemble in ensembles
                if orchestra.ID == ensemble._orchestra_ID]

many_to_many = [(conductor._name, conductor._work_experience, orchestra_name)
                for orchestra_name, _, conductorID in many_to_many_temp
                for conductor in conductors if conductor._conductorID == conductorID]

#Task1: Выведите список всех всязанных дирижеров и оркестров, отсортированных по дирижерам:

result1 = [(conductor_name, orchestra_name)
           for conductor_name, _, orchestra_name, _ in one_to_many]

result1 = sorted(result1, key=lambda x: x[0])

for i in result1:
    print(i)

print()
#Task2: Выведите список отделов, дирижеры которых имеют стаж работы меньше 10 лет, отсортировать по стажу работы

result2 = [(orchestra_name, conductor_name, work_experience)
           for conductor_name, work_experience, orchestra_name, _ in one_to_many
           if work_experience <= 10]

result2 = sorted(result2, key=lambda x: x[2])

for i in result2:
    print(i)

print()
#Task3:

result3 = [(orchestra_name, conductor_name)
           for conductor_name, _, orchestra_name in many_to_many]

result3 = set(sorted(result3, key=lambda x: x[0]))

for i in result3:
    print(i)
